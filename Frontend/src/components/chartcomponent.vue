<template>
  <div>
    <v-row class="d-flex justify-center align-center mt-8 mb-4">

    
    <div >
      <h1>Monthly Sales</h1>
      <apexchart width="500" type="line" :options="options" :series="series"></apexchart>
    </div>
  </v-row>
   
  </div>
    </template>

<script>

import apexchart from 'vue-apexcharts'
import Eventservice  from '../service/Eventservice';

export default{
    data(){
        return{
            options: {
        chart: {
          id: 'vuechart-example'
        },
        xaxis: {
          categories: ["jan", "feb", "Mar", "Apr", "May", "June", "July","Aug","Sep","Oct","Nov","Dec"]
        }
      },
      series: [{
        name: 'series-1',
        data: [ ]
      }],
   
    }

        },

        components:{
            apexchart
        },

        mounted(){
          Eventservice.FetchMonthSales()
        .then((response) => {
          if(response.data.status =="S"){
            const salesData = response.data.month_obj;
          const formattedData = [
            salesData.jan,
            salesData.feb,
            salesData.mar,
            salesData.apr,
            salesData.may,
            salesData.jun,
            salesData.jul,
            salesData.aug,
            salesData.sep,
            salesData.oct,
            salesData.nov,
            salesData.dec
          ];

          this.series = [{
            name: 'Monthly Sales',
            data: formattedData
          }];
            
            // this.series[0].data = response.data.month_obj
         
 
            // const monthArr = response.data.month_arr[0];
            // this.series =[{name: 'Sales',data: Object.values(monthArr)}]
            // console.log(this.series[0].data ,"Series");
            // console.log("weekly",response.data.week_arr[0]); 
     
            
            
            // console.log("sERI",this.series[0].data)
          };
        })
        .catch((error) => {
          console.log(error);
        });


        }
    }

</script>