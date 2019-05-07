<template>
  <div id="app">
    <nav>
      <div class="nav-wrapper orange darken-1">
        <a href="#" class="brand-logo center">
          QUEM DEVE SER
          <b>ELIMINADO</b>?
        </a>
      </div>
    </nav>

    <div class="container row center">
      <form @submit.prevent="check">
        <div class="center row">
          <div class="columnHome row">
            <h4 class="name">Maria</h4>
          </div>
          <div class="columnHome row">
            <h4 class="name">João</h4>
          </div>
        </div>
        <hr>
        <div class="center row">
          <div class="columnHome">
            <label>
              <input type="radio" id="one" value="Maria" v-model="picked">
              <img src="../assets/images/sister.png" class="container center">
            </label>
          </div>
          <div class="columnHome">
            <label>
              <input type="radio" id="two" value="Joao" v-model="picked">
              <img src="../assets/images/brother.png" class="container center">
            </label>
          </div>
        </div>
        <hr>
        <div class="center row">
          <div class="columnHome">
            <p class="container row center">
              Para eliminar a
              <b>Sister</b> pelo telefone disque
              <b>0800-123-001</b> ou mande um SMS para
              <b>8001</b>
            </p>
          </div>
          <div class="columnHome">
            <p class="container row center">
              Para eliminar o
              <b>Brother</b> pelo telefone disque
              <b>0800-123-002</b> ou mande um SMS para
              <b>8002</b>
            </p>
          </div>
        </div>
        <div id="Footer" class="center jumbotron">
          <button class="waves-effect waves-light btn-large orange bt" :disabled="!checked || !picked" >ENVIE SEU VOTO AGORA</button>
        </div>
      </form>
      <vue-recaptcha class="waves-effect captcha"
        @verify="onVerify"
        @expired="onExpired"
        sitekey="6LdunZ8UAAAAAP5gz-M4eq5duwfZsPPuODcezVfS"
      ></vue-recaptcha>
    </div>
  </div>
</template>

<script>
import Voting from "../services/voting";
import VueRecaptcha from "vue-recaptcha";
export default {
  components: { VueRecaptcha },
  name: "Home",
  data() {
    return {
      picked: "",
      checked: "",
      participant: {
        id: "",
        name: "",
        votes: 0,
        elapsed_hours: 0
      }
    };
  },

  methods: {
    onSubmit: function() {
      this.$refs.invisibleRecaptcha.execute();
    },
    onVerify: function() {      
      this.checked = "ok"
    },
    onExpired: function() {
      this.checked = ""
    },
    check() {
      if (this.picked) {
        this.participant.name = this.picked;
        Voting.Vote(this.participant)
          .then(res => {
            if (res.status === 200 || res.status === 201) {
              this.$router.push({
                name: "result",
                params: { Pid: this.picked.toUpperCase() }
              });
            } else {
              alert("Server offline: Aguarde um momento e tente novamente");
            }
          })
          .catch(e => {
            alert(e + ": Aguarde um momento e tente novamente");
          });
      } else {
        alert("Esolha um participante para se eliminado");
      }
    }
  }
};
</script>

<style>
.bt {
  font-size: 26px;
}
.columnHome {
  width: 50%;
  float: left;
  padding: 8px;
}

.captcha {
  padding: 8px;
  border-radius: 5px;
  line-height: 54px;
  font-size: 15px;
}
</style>