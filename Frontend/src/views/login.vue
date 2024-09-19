<!-- <template>
  <v-app>
    <v-container>
      <v-row justify="center">
        <v-col cols="12" sm="8" md="6">
          <v-card class="elevation-10">
            <v-card-title class="blue darken-1 white--text mb-8">Login</v-card-title>
            <v-card-text>
              <v-form @submit.prevent="login">
                <v-text-field v-model="userId" label="User ID" outlined></v-text-field>
                <v-text-field v-model="password" label="Password" outlined type="password"></v-text-field>
                <v-btn
      dark
      @click="login()"
    >
      Login
    </v-btn>
    <v-snackbar
      v-model="snackbar"
    >
      {{ text }}

      <template v-slot:action="{ attrs }">
        <v-btn
          color="pink"
          text
          v-bind="attrs"
          @click="snackbar = false"
        >
          Close
        </v-btn>
      </template>
    </v-snackbar>

                
              </v-form>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </v-app>
</template> -->
<template>
  <v-app>
    <v-container>
      <v-row justify="center" align="center">
        <v-col cols="12" sm="8" md="6">
          <v-card class="elevation-10">
<v-card-title class="blue darken-1 white--text mb-8 black--text" style="background-color: black"> <v-icon>mdi-hospital</v-icon>MEDAPP</v-card-title>
            <v-card-text>
            
                <v-text-field
                  v-model="userId"
                  label="User ID"
                  outlined
                  :rules="userIdRules"
                  autocomplete="User ID"
                ></v-text-field>
                <v-text-field
                  v-model="password"
                  label="Password"
                  outlined
                  type="password"
                  :rules="passwordRules"
                  autocomplete="Password"
                ></v-text-field>

                <v-btn  class="blue darken-1 white--text mb-8 black--text" @click="login">Login</v-btn>
             
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>

  
    </v-container>
  </v-app>
</template>

<script>
import EventService from "../service/Eventservice";

export default {
  data() {
    return {
      userId: "",
      password: "",
      error: false,
      currentuser: "",
      snackbar: false,
      text: "",
      historyid: 0,
      role: "",
      userIdRules: [(v) => !!v || "User ID is required"],
      passwordRules: [(v) => !!v || "Password is required"],
      responseData: [],
      login_id:0,

    };
  },
  methods: {
    login() {
      if (this.userId == "" || this.password == "") {
        this.snackbar = true;
        this.text = "please fill the fields";

        return;
      } else {

        const struct = {
          user_id: this.userId.trim(),
          password: this.password,
        };

        EventService.sendLogin(struct)
          .then((response) => {
            // console.log("Login Response:", response);
            const user_details=response.data.user_details
            const { user_id, password, role } = response.data.user_details;
            if( user_id !== "" || password !== "" || role !== ""){

          //  console.log("login res",response)
            // if (response.data.user_details != null) {
            //   const user_details = response.data.user_details;
            //   this.login_id=response.data.user_details.login_id
              // console.log("loginn",this.login_id)

              const historystruct = {
                login_id: user_details.login_id,
              };

              EventService.sendLoginHistory(historystruct)
                .then((response) => {
                  this.historyid = response.data.login_history_id;

                  // console.log("Login History Response:", response);

                  // Check if user_details and its properties exist before using them
                  if (
                    user_details &&
                    user_details.role &&
                    user_details.login_id
                  ) {

                   // Set cookies
                    document.cookie = `user_id=${user_details.user_id};`;
                    document.cookie = `role=${user_details.role};`;
                    document.cookie = `login_id=${user_details.login_id};`; 

                    // console.log("document",document.cookie.length)

                    // console.log("Role:", user_details.role);
                    // console.log("Login ID:", user_details.login_id);  

                  

                    this.$router.push({
                      name: "Dashboard",
                      params: {
                        user_id:user_details.user_id,
                        role: user_details.role,
                        login_id: user_details.login_id,
                        login_history_id: this.historyid,
                      },
                    });
                  } else {
                    // console.log("Role or login ID not found in user details");
                    // Handle this case as needed
                  }
                })
                .catch((error) => {
                  console.log("History Error:", error);
                });
            } else {
              alert("User details are not found");
              this.userId=" "
              this.password=" "
              // Handle this case as needed
            }
          })
          .catch((error) => {
            console.log("Login Error:", error);
          });
      }
    },
  },
};
</script>

<style scoped>
.v-card {
  margin-top: 100px;
}
</style>
