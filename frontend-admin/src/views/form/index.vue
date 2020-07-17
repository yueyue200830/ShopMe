<template>
  <div class="app-container">
    <el-card class="box-card">
      <div slot="header">
        <span>个人信息</span>
      </div>
      <el-form ref="form" :model="form" :rules="rules" status-icon label-width="100px">
        <el-form-item label="用户名">
          <el-input v-model="form.name" disabled />
        </el-form-item>
        <el-form-item label="原密码" prop="originalPassword">
          <el-input
            v-model="form.originalPassword"
            show-password
            type="password"
            placeholder="请输入原密码"
            maxlength="16"
            minlength="6"
          />
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input
            v-model="form.password"
            show-password
            type="password"
            placeholder="请输入新密码"
            maxlength="16"
            minlength="6"
          />
        </el-form-item>
        <el-form-item label="确认密码" prop="confirmPassword">
          <el-input
            v-model="form.confirmPassword"
            show-password
            type="password"
            placeholder="请确认密码"
            maxlength="16"
            minlength="6"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">确认修改</el-button>
          <el-button @click="resetForm('registerForm')">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

  </div>
</template>

<script>
import { updateManager } from '@/api/manager'

export default {
  data() {
    const validateOriginalPassword = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入原密码'))
      } else {
        callback()
      }
    }
    const checkPassword = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码'))
      } else {
        const passwordVerifier = /^.*(?=.{6,16})(?=.*[a-zA-Z]).*$/
        if (passwordVerifier.test(value)) {
          callback()
        } else {
          callback(new Error('输入6-16位密码，需包含英文字母和数字'))
        }
      }
    }
    const validatePassword = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'))
      } else if (value !== this.form.password) {
        callback(new Error('两次输入密码不一致!'))
      } else {
        callback()
      }
    }
    return {
      form: {
        name: '',
        password: '',
        originalPassword: '',
        confirmPassword: '',
        id: undefined,
      },
      rules: {
        originalPassword: [
          { validator: validateOriginalPassword, trigger: 'blur' }
        ],
        password: [
          { validator: checkPassword, trigger: ['blur', 'change'] }
        ],
        confirmPassword: [
          { validator: validatePassword, trigger: 'blur' }
        ],
      }
    }
  },
  created() {
    this.form.name = this.$store.getters.name
    this.form.id = this.$store.getters.token
  },
  methods: {
    onSubmit() {
      this.$refs['form'].validate(valid => {
        if (!valid) {
          return
        }
        const data = new FormData()
        data.append('id', this.form.id)
        data.append('oldPassword', this.form.originalPassword)
        data.append('newPassword', this.form.password)
        updateManager(data).then(response => {
          if (response !== 0) {
            this.$message.error('更新密码失败，请重试')
          } else {
            this.$message.success('密码更新成功')
            this.resetForm('form')
          }
        })
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
      this.form.name = this.$store.getters.name
      this.form.id = this.$store.getters.token
    },
  }
}
</script>

<style scoped>
  .app-container{
    display: flex;
  }

  .box-card {
    width: 550px;
    margin: 40px auto;
  }
</style>

