<template>
  <div>
    <v-container style="background: linear-gradient(to bottom right, #2196F3, #1976D2);">
      <v-row>
        <v-col cols="12">
          <v-expansion-panels>
            <v-expansion-panel>
              <v-expansion-panel-header>Add User</v-expansion-panel-header>
              <v-expansion-panel-content>
                <v-container>

                  <v-form @submit.prevent="adduser">
                    <v-row>
                      <v-col cols="12" sm="4">
                        <v-text-field
                          v-model="userId"
                          label="User ID"
                          autocomplete="User ID"
                        ></v-text-field>
                      </v-col>
                      <v-col cols="12" sm="4">
                        <v-text-field
                          v-model="password"
                          label="Password"
                          type="password"
                          autocomplete="Password"
                        ></v-text-field>
                      </v-col>
                      <v-col cols="12" sm="4">
                        <v-select
                          v-model="role"
                          :items="apiroles"
                          label="Role"
                        ></v-select>
                      </v-col>
                      <v-col cols="12">
                        <v-btn color="primary" type="submit">Add User</v-btn>
                      </v-col>
                    </v-row>
                  </v-form>
                  
                </v-container>
              </v-expansion-panel-content>
            </v-expansion-panel>
          </v-expansion-panels>

          <v-snackbar v-model="snackbar" :timeout="3000"  :color="snackbarColor" bottom center>
            {{ snackbarMessage }}
            <v-btn color="black" text @click="snackbar = false">Close</v-btn>
          </v-snackbar>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>
<script>
import EventService from "../service/Eventservice";
export default {
  data() {
    return {
      userId: "",
      password: "",
      role: null,
      roles: ["Biller", "Manager", "Systemadmin", "Inventry"],
      login: false,
      apiroles: [],
      userroles: [],
      snackbar: false,
      snackbarMessage: "",
      snackbarColor:''
    };
  },

  created() {
    EventService.SendAddUserDrop()
      .then((response) => {
        // console.log(response);
        this.apiroles = response.data.role;
        // console.log(this.apiroles);
      })
      .catch((error) => {
        console.log(error);
      });
  },


  methods: {
    adduser() {
      if (!this.userId || !this.password || !this.role) {
        this.snackbarMessage = " Please fill the all fields ";
            this.snackbar = true;
            this.snackbarColor="red"
            return;
      }
      if (this.userId.includes(" ") || this.password.includes(" ")) {
        this.snackbarMessage = " Cannot Contain Space ";
            this.snackbar = true;
            this.snackbarColor="red"
            return;
      }
      if (/\d/.test(this.userId)) {
        this.snackbarMessage = "Userid Cannot Contain numbers ";
            this.snackbar = true;
            this.snackbarColor="red"
            return;
      }

      // console.log(this.userId);

      const user = {
        user_id: this.userId,
        password: this.password,
        role: this.role,
      };
      EventService.FetchUsers(user)
        .then((response) => {
          console.log(response)

          if (response.data.status == "E") {
            this.snackbarMessage = "User Exist ";
            this.snackbar = true;
            this.snackbarColor="red"
            return;
          } else if(response.data.status == "S") {
            this.snackbarMessage = "New User Added Successfully ";
            this.snackbar = true;
            this.snackbarColor="green"
            return;
          }else{
            this.snackbarMessage = "Internal Error";
            this.snackbar = true;
            this.snackbarColor="red"
          }
        })
        .catch((error) => {
          console.log(error);
        });

    

      this.login = true;
      (this.userId = ""), (this.password = ""), (this.role = null);
    },
  },
};
</script>
