<template>
  <div>
    <nav>
      <div class="nav-wrapper orange darken-1">
        <a href="#" class="brand-logo center">
          QUEM DEVE SER
          <b>ELIMINADO</b>?
        </a>
      </div>
    </nav>

    <div>
      <div class="row">
        <div class="center">
          <h4 class="center">
            PARABÉNS! O SEU VOTO PARA ELIMINAR
            <b>{{proId}}</b>
            FOI ENVIADO COM SUCESSO.
          </h4>
        </div>
      </div>
      <div class="center row">
        <div class="column">
          <label>
            <img src="../assets/images/sister.png" class="center">
          </label>
        </div>
        <div class="column">
          <section v-if="errored">
            <p>Pedimos desculpas, não estamos conseguindo recuperar as informações no momento. Por favor, tente novamente mais tarde.</p>
          </section>
          <section v-else>
            <div v-if="loading">Carregando...</div>
            <div v-else>
              <Donut
                background="white"
                foreground="grey"
                :size="300"
                unit="px"
                :thickness="20"
                has-legend
                legend-placement="bottom"
                :sections="sections"
                :total="totaldonut"
              >{{"Total: "+totaldonut}}</Donut>
            </div>
          </section>
        </div>
        <div class="column">
          <label>
            <img src="../assets/images/brother.png" class="center">
          </label>
        </div>
      </div>
    </div>
    <label>{{teste}}</label>
  </div>
</template>
<style>
.header {
  margin-top: 32px;
}

.name {
  width: 100%;
  font-size: 60px;
  margin-left: auto;
  margin-right: auto;
  text-align: center;
}
.center {
  width: 100%;
  font-size: 24px;
  margin-left: auto;
  margin-right: auto;
  text-align: center;
}

.row {
  padding: 5px;
}

.column {
  width: 30%;
  float: left;
  padding: 8px;
}

label {
  width: 100%;
}

label > input {
  visibility: hidden;
  position: absolute;
}

label > input + img {
  cursor: pointer;
  border: 20px solid transparent;
}

label > input:checked + img {
  border: 4px solid orange;
}
</style>
<script>
import Voting from "../services/voting";
import Donut from "./donut/Donut.vue";
import colors from "./donut/utils/colors";
import "./donut/styles/normalize.css";
import "./donut/styles/site.css";

export default {
  components: { Donut },
  name: "App",
  data() {
    return {
      loading: true,
      errored: false,
      proId:this.$route.params.Pid,
      participant: {
        Name: "",
        votes: ""
      },
      participants: [],
      errors: [],
      totaldonut: 0,
      sections: [
        { label: "", value: 20, color: colors[2] },
        { label: "", value: 30, color: colors[3] }
      ]
    };
  },
  created() {
    this.fetch();
  },
  methods: {
    async fetch() {
      await Voting.Get()
        .then(response => {
          if (response.status === 200 || response.status === 201) {
            this.participants = response.data.participants
            
            this.sections[0].label =this.participants[0].Name
            this.sections[0].value =parseInt(this.participants[0].Votes)
            
            this.sections[1].label = this.participants[1].Name
            this.sections[1].value =parseInt(this.participants[1].Votes)

            this.totaldonut = parseInt(this.participants[0].Votes) + parseInt(this.participants[1].Votes)
            
          } else {
            alert("Erro ao buscar relação de votos");
            this.errored = true;
          }
        })
        .catch(e => {
          alert(e);
        })
        .finally(() => (this.loading = false));
    }
  }
};
</script>

