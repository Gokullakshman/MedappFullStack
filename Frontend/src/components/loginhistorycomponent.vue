<template>
  <div>


  <!-- <v-expansion-panels>
    <v-expansion-panel title="Login History">
    
      <template v-slot:default>
        <v-simple-table>
          <template v-slot:default>
            <thead>
              <tr>
                <th class="text-center">User ID</th>
                <th class="text-center">Login Time</th>
                <th class="text-center">Login Date</th>
                <th class="text-center">Logout Time</th>
                <th></th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(entry, index) in history" :key="index">
                <td class="text-center">{{ entry.user_id }}</td>
                <td class="text-center">{{ entry.login_time }}</td>
                <td class="text-center">{{ entry.login_date }}</td>
                <td class="text-center">{{ entry.logout_time }}</td>
              </tr>
            </tbody>
          </template>
        </v-simple-table>
      </template>
    </v-expansion-panel>
  </v-expansion-panels> -->
  <v-card>
    <v-card-title>
      <v-text-field
        v-model="search"
        append-icon="mdi-magnify"
        label="Search"
        single-line
        hide-details
      ></v-text-field>
    </v-card-title>
    <v-data-table
      :headers="headers"
      :items="history"
      :search="search"
    ></v-data-table>
  </v-card>
</div>
</template>
<!-- <v-data-table
:headers="headers"
:items="stock"
:search="search"
></v-data-table> -->

<script>
import EventService from "../service/Eventservice";
export default {
  data() {
    return {
      history: [],
      search: "",
      headers: [
        { text: "UserId", value: "user_id" },
        { text: "LoginTime", value: "login_time" },
        { text: "LoginDate", value: "login_date" },
        { text: "LogoutTime", value: "logout_time" },
      ],
    };
  },

  created() {
    EventService.sendHistory()
      .then((response) => {
        // console.log(response);
        this.history = response.data.history_details_arr;
      })
      .catch((error) => {
        console.log(error);
      });
  },
};
</script>
