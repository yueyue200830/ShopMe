<template>
  <div>
    <el-header>
      <div class="head-title" @click="toMainPage">ME</div>
    </el-header>
    <el-main class="user-body">
      <div class="user-title">
        用户注册
      </div>
      <el-form
          :model="registerForm"
          status-icon
          :rules="rules"
          ref="registerForm"
          class="register-form"
          label-width="100px">
        <el-form-item label="用户名" prop="name">
          <el-input
              type="name"
              v-model="registerForm.name"
              maxlength="20"
              minlength="4"
              placeholder="请输入用户名"
          />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input
              type="email"
              v-model="registerForm.email"
              maxlength="45"
              minlength="4"
              placeholder="请输入邮箱"
          />
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input
              type="password"
              placeholder="请输入密码"
              show-password
              maxlength="16"
              minlength="6"
              v-model="registerForm.password"
          />
        </el-form-item>
        <el-form-item label="确认密码" prop="confirmPassword">
          <el-input
              type="password"
              placeholder="请确认密码"
              show-password
              maxlength="16"
              minlength="6"
              v-model="registerForm.confirmPassword"
          />
        </el-form-item>
        <el-form-item class="register-button">
          <el-button class="button-margin" type="primary" @click="register" :loading="signingUp">注册</el-button>
          <el-button class="button-margin" @click="resetForm('registerForm')">重置</el-button>
          <el-button type="text" @click="login">登录</el-button>
        </el-form-item>
      </el-form>
    </el-main>
    <user-footer/>
  </div>
</template>

<script>
  import UserFooter from '../components/UserFooter';

  export default {
    name: 'UserRegister',
    components: {UserFooter},
    data() {
      const checkName = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请输入用户名'))
        } else {
          let nameVerifier = /^[a-zA-Z0-9_\u4e00-\u9fa5]{4,20}$/
          if (nameVerifier.test(value)) {
            this.$http
              .get('/api/checkUserNameExist', {
                params: {
                  name: value
                }
              })
              .then(response => {
                if (response.data === 0) {
                  callback()
                } else {
                  callback(new Error('用户名已存在'))
                }
              })
              .catch(error => {
                console.log(error)
                this.$message.error('验证用户名失败')
              })
          } else {
            callback(new Error('用户名不合法，请输入4-20个字符，支持大小写中文'))
          }
        }
      }
      const checkEmail = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请输入邮箱'))
        } else {
          let emailVerifier = /^[a-zA-Z0-9]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/
          if (emailVerifier.test(value)) {
            this.$http
              .get('api/checkUserEmailExist', {
                params: {
                  email: value
                }
              })
              .then(response => {
                if (response.data === 0) {
                  callback()
                } else {
                  callback(new Error('邮箱已注册，请直接登录'))
                }
              })
              .catch(error => {
                console.log(error)
                this.$message.error('验证邮箱失败')
              })
          } else {
            callback(new Error('请输入正确的邮箱'))
          }
        }
      }
      const checkPassword = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请输入密码'))
        } else {
          let passwordVerifier = /^.*(?=.{6,16})(?=.*[a-zA-Z]).*$/
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
        } else if (value !== this.registerForm.password) {
          callback(new Error('两次输入密码不一致!'))
        } else {
          callback()
        }
      }
      return {
        signingUp: false,
        registerForm: {
          name: '',
          password: '',
          confirmPassword: '',
          email: ''
        },
        rules: {
          name: [
            {validator: checkName, trigger: ['blur', 'change'], required: true}
          ],
          email: [
            {validator: checkEmail, trigger: ['blur', 'change'], required: true}
          ],
          password: [
            {validator: checkPassword, trigger: ['blur', 'change'], required: true}
          ],
          confirmPassword: [
            {validator: validatePassword, trigger: 'blur', required: true}
          ],
        }
      }
    },
    methods: {
      register() {
        this.signingUp = true
        this.$refs['registerForm'].validate((valid) => {
          if (valid) {
            this.$http
              .post('/api/register', this.registerForm)
              .then(response => {
                if (response.data === 0) {
                  this.$message.success('注册成功，请登录！')
                  this.$router.push('/login')
                } else if (response.data === 1) {
                  this.$message.error('用户名已存在，请重试！')
                } else if (response.data === 2) {
                  this.$message.error('邮箱已存在，请重试！')
                } else {
                  this.$message.error('注册失败，请重试！')
                }
              })
              .catch(error => {
                console.log(error)
                this.$message.error('注册失败，请重试！')
              })
              .finally(() => {
                this.signingUp = false
              })
          } else {
            this.signingUp = false
          }
        })
      },
      login() {
        this.$router.push('/login')
      },
      resetForm(formName) {
        this.$refs[formName].resetFields()
      },
      toMainPage() {
        this.$router.push('/')
      }
    }
  }
</script>

<style scoped>
  .el-header {
    background-color: #333;
    color: white;
    text-align: center;
    line-height: 60px;
    display: flex;
    flex-direction: row;
    font-size: 20px;
  }

  .head-title {
    margin: 0 auto;
    padding: 0 40px;
  }

  .head-title:hover {
    cursor: pointer;
  }

  .user-body {
    min-height: calc(100vh - 120px);
  }

  .user-title {
    text-align: center;
    font-size: 30px;
    margin: 100px 0 40px;
    color: #303133;
  }

  .register-form {
    width: 400px;
    margin: 10px auto;
  }

  .register-button {
    margin-top: 40px;
  }
</style>