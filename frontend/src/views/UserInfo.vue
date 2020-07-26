<template>
  <user-product-component>
    <div class="user-title">
      个人信息
    </div>
    <div class="user-info-body">
      <div class="user-avatar">
        <el-image
            style="width: 200px; height: 200px"
            :src="infoForm.avatar"
            fit="fit"/>
      </div>
      <div class="info-form">
        <el-form
            :model="infoForm"
            status-icon
            :rules="rules"
            ref="infoForm"
            label-width="100px">
          <el-form-item label="用户名" prop="name">
            <el-input
                type="name"
                v-model="infoForm.name"
                maxlength="20"
                minlength="4"
                placeholder="请输入新用户名"
            />
          </el-form-item>
          <el-form-item label="邮箱" prop="email">
            <el-input
                type="email"
                v-model="infoForm.email"
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
  import UserProductComponent from '../components/UserProductComponent'
  import { mapState } from 'vuex'

  export default {
    name: 'UserInfo',
    components: {UserProductComponent},
    computed: {
      hasLoggedIn() {
        return this.$store.getters.hasLoggedIn
      },
      userID() {
        return this.$store.getters.getUserID
      },
      ...mapState([
         'Token',
       ])
    },
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
                  name: value,
                  id: this.userID
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
              .get('/api/checkUserEmailExist', {
                params: {
                  email: value,
                  id: this.userID
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
        infoForm: {
          name: '',
          email: '',
          avatar: '',
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
    created() {
      this.getInfo()
    },
    methods: {
      updateInfo() {
        this.$refs['infoForm'].validate(valid => {
          if (valid) {
            this.$http
              .put('/api/user', {
                id: this.userID,
                name: this.infoForm.name,
                email: this.infoForm.email
              }, {
                headers: {
                  Bearer: this.Token == null ? "" : this.Token
                }
              })
              .then(response => {
                if (response.data !== 0) {
                  this.$message.error('更新信息失败')
                } else {
                  this.$message.success('个人信息更新成功')
                  this.getInfo()
                }
              })
          }
        })
      },
      getInfo() {
        this.$http
          .get('/api/user', {
            headers: {
              Bearer: this.Token == null ? "" : this.Token
            },
            params: {
              id: this.userID
            }
          })
          .then(response => {
            if (response.data.code !== 0) {
              this.$message.error('获取个人信息失败')
            } else {
              let info = response.data.data
              info.avatar = '/api/avatar/' + info.avatar
              this.infoForm = info
            }
          })
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

  .user-avatar {
    margin: 50px 50px 50px 100px;
  }

  .info-form {
    margin: 60px auto 40px 40px;
  }
</style>