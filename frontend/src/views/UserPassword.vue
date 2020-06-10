<template>
  <user-product-component>
    <div class="user-title">
      修改密码
    </div>
    <div class="user-info-body">
      <div class="info-form">
        <el-form
            :model="updateForm"
            status-icon
            :rules="rules"
            ref="updateForm"
            label-width="100px">
          <el-form-item
              label="旧密码"
              :rules="[{ required: true, message: '密码不可为空'}]"
              prop="password">
            <el-input
                type="password"
                placeholder="请输入旧密码"
                v-model="updateForm.originalPassword"
                show-password
                maxlength="18"
            />
          </el-form-item>
          <el-form-item label="新密码" prop="password">
            <el-input
                type="password"
                placeholder="请输入新密码"
                show-password
                maxlength="16"
                minlength="6"
                v-model="updateForm.password"
            />
          </el-form-item>
          <el-form-item label="确认新密码" prop="confirmPassword">
            <el-input
                type="password"
                placeholder="请确认新密码"
                show-password
                maxlength="16"
                minlength="6"
                v-model="updateForm.confirmPassword"
            />
          </el-form-item>
          <el-form-item class="register-button">
            <el-button
                class="button-margin"
                type="primary"
                @click="updatePassword"
                :loading="updating">
              更新密码
            </el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
  </user-product-component>
</template>

<script>
  import UserProductComponent from '../components/UserProductComponent';
  export default {
    name: 'UserPassword',
    components: {UserProductComponent},
    data() {
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
        updating: false,
        updateForm: {
          password: '',
          confirmPassword: '',
          originalPassword: '',
        },
        rules: {
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
      updatePassword() {

      }
    }
  }
</script>

<style scoped>
  .user-title {
    background-color: white;
    text-align: center;
    font-size: 24px;
    padding: 40px 0 0;
    margin-top: 20px;
  }

  .user-info-body {
    background-color: white;
    display: flex;
    padding: 40px 150px;
  }

  .info-form {
    margin: 60px auto 40px;
  }
</style>