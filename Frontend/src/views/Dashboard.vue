<template>
  <div id="app">
    <div v-if="role === 'Biller'">
      <Billercomponent
        :login_history_id="login_history_id"
        :login_id="login_id"
        :sales="sales"
        :user_id="user_id"
        :flag="flag"
      ></Billercomponent>
    </div>

    <div v-else-if="role === 'Manager'">
      <Managercomponent :login_history_id="login_history_id"></Managercomponent>
    </div>

    <div v-else-if="role === 'System Admin'">
      <Systemadmin :login_history_id="login_history_id"></Systemadmin>
    </div>

    <div v-else-if="role === 'Inventry'">
      <Inventry :login_history_id="login_history_id"></Inventry>
    </div>
  </div>
</template>

<script>
import EventService from "../service/Eventservice";
import Billercomponent from "../components/billerscomponent.vue";
import Managercomponent from "../components/managercomponent.vue";
import Systemadmin from "../components/systemadmincomponent.vue";
import Inventry from "../components/inventrycomponent.vue";

export default {
  data() {
    return {
      role: "",
      login_id: 0,
      login_history_id: 0,
      user_id: "",
      sales: 0,
      yesterdaysales: 0,
      per:0,
      flag:-1 ,
      duserid:'',
      drole:''
    };
  },

  created() {
    this.login_history_id = this.$route.params.login_history_id;

    this.role = this.$route.params.role;
    this.user_id = this.$route.params.user_id;
    // this.user_id=this.$store.state.go_user_id
    this.login_id = this.$route.params.login_id;
    this.$store.state.go_login_id = this.login_id;

    // console.log("login  id dashboard ", this.login_id);

    this.summa(); 


     // Get cookies when the component is created
     this.duserid = this.getCookie("user_id");
    this.drole = this.getCookie("role"); 
    

    // console.log("document user_id",this.duserid)
    // console.log("document role",this.drole)
  },
  methods: {
    summa() {
      const struct = {
        user_id: this.user_id,
      };

      EventService.sendToday(struct)
        .then((response) => {
          // console.log(response);
          this.sales = Math.round(response.data.today_sales);
          this.yesterdaysales = Math.round(response.data.yesterday_sales);
          if (this.sales == 0 && this.yesterdaysales == 0) {
            this.per = 0;
          } else if (this.sales != 0 && this.yesterdaysales == 0) {
            this.per = 100;
          } else if (this.sales == 0 && this.yesterdaysales != 0) {
            this.per = 0;
          } else {
            if (this.sales > this.yesterdaysales) {
              this.per =
                Math.abs(
                  (this.sales - this.yesterdaysales) / this.yesterdaysales) * 100;
            } else {
              this.per = 0; // or any other appropriate value if you want to handle when today's sales are not greater
            }
          }

          // console.log(this.sales);
          // console.log(this.sales)
          // console.log(this.yesterdaysales)
          // console.log(Math.abs(Math.round(this.per)))  

          if (typeof this.per === "number" && !isNaN(this.per)) {
        if (this.per > 0) {
          this.flag=1
        } else if (this.per < 0) {

          this.flag=0
        } 
      }
})
        .catch((error) => {
          console.log(error);
          this.sales = 0;
        });
    },

     getCookie(name) {
     
      const cookies = document.cookie.split(';');
      
      for (let i = 0; i < cookies.length; i++) {
        const cookie = cookies[i].trim();
     
        if (cookie.startsWith(name + '=')) {
        
          return cookie.substring(name.length + 1); 

        }
      }
      
      return "";
    }
  },

  components: {
    Billercomponent,
    Managercomponent,
    Systemadmin,
    Inventry,
  },
};
</script>
