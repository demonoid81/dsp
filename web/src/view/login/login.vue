<template>
  <div class="login">
    <div class="login-con">
      <Card icon="log-in" title="Авторизация" :bordered="false">
        <div class="form-con">
          <login-form @on-success-valid="handleSubmit"></login-form>
          <p class="login-tip">Введите имя пользователя и пароль</p>
        </div>
      </Card>
    </div>
  </div>
</template>

<script>
import LoginForm from 'components/login-form'
import { mapActions } from 'vuex'
export default {
  name: "login",
  components: {
    LoginForm
  },
  methods: {
    ...mapActions({
      handleLogin: 'user/handleLogin'
    }),
    async handleSubmit ({ username, password }) {
      this.handleLogin({ username, password }).then(res => {
        console.log(res)
        // this.getUserInfo().then(res => {
          this.$router.push({
            name: this.$config.homeName
          })
        // })
      })
    }
  }
}
</script>

<style lang="less">
@import 'login.less';
</style>