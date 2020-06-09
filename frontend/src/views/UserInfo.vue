<template>
  <user-product-component>
    <div class="user-info-body">
      <div class="user-avatar">
        <el-image
            style="width: 200px; height: 200px"
            :src="url"
            :fit="fit"></el-image>
      </div>
      <div class="info-form">
        <el-form :model="registerForm" status-icon :rules="rules" ref="registerForm" label-width="100px"
                 class="register-form">
          <el-form-item label="用户名" prop="name">
            <el-input
                type="name"
                v-model="registerForm.name"
                maxlength="20"
                minlength="4"
                placeholder="请输入新用户名"
            />
          </el-form-item>
          <el-form-item label="邮箱" prop="email">
            <el-input
                type="email"
                v-model="registerForm.email"
                maxlength="45"
                minlength="4"
                placeholder="请输入新邮箱"
            />
          </el-form-item>
          <el-form-item class="register-button">
            <el-button
                class="button-margin"
                type="primary"
                @click="updateInfo"
                :loading="updating">
              更新个人信息
            </el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
  </user-product-component>
</template>

<script>
  import UserProductComponent from "../components/UserProductComponent";

  export default {
    name: 'UserInfo',
    components: {UserProductComponent},
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
                  callback(new Error('邮箱已被注册，请更换账号登录'))
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
      return {
        updating: false,
        registerForm: {
          name: '',
          email: ''
        },
        rules: {
          name: [
            {validator: checkName, trigger: ['blur', 'change'], required: true}
          ],
          email: [
            {validator: checkEmail, trigger: ['blur', 'change'], required: true}
          ],
        }
      }
    },
    methods: {
      updateInfo() {

      }
    }
  }
</script>

<style scoped>
  .user-info-body {
    background-color: white;
    display: flex;
    padding: 40px 150px;
  }

  .user-avatar {
    margin: 50px 50px 50px 100px;
  }

  .info-form {
    margin: 60px auto 40px 40px;
  }
</style>