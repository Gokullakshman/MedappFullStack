<template>
  <div>
    <v-container  class="d-flex flex-column justify-center align-center">

   <v-row class="mb-10">

   
    <div>
      <h1>Weekly Sales</h1>
      <apexchart
        width="500"
        type="line"
        :options="options"
        :series="series"
      ></apexchart>
    </div>
  </v-row>
  <v-row class="mb-10">
    <div>
      <h1>User's Sales</h1>
      <apexchart
        width="380"
        type="donut"
        :options="options1"
        :series="series1"
      ></apexchart>
    </div>
  </v-row>
  <v-row class="mb-10">
    <div>
      <div>
        <h1>This month sales</h1>
  <apexchart  width="500" type="bar" :options="options2" :series="series2"></apexchart>
</div>

    </div>
  </v-row>
  </v-container>
  </div>
</template>

<script>
import apexchart from "vue-apexcharts";
import Eventservice from "../service/Eventservice";
// 

export default {
  data() {
    return {
      biller: [],
      dalilySale: [],
      WeeklyArr:null,
      options: {
        chart: {
          id: "vuechart-example",
        },
        xaxis: {
          categories: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
        },
      },
      series: [
        {
          name: "series-1",
          data: [],
        },
      ],


      options1: {},
      series1: [],


      options2: {
        chart2: {
          id: 'vuechart-example'
        },
        xaxis2: {
          categories: []
        }
      },
      series2: [{
        name: 'series-1',
        data: []
      }]
    };

    },
    


  components: {
    apexchart,
  },

  mounted() {
    Eventservice.FetchWeeklySales()
      .then((response) => {
        if (response.data.status == "S") {
          console.log(response.data)
          const weeksales = response.data.week_obj;
          const salesdate=[weeksales.mon,weeksales.tue,weeksales.wed,weeksales.thu,weeksales.fri,weeksales.sat,weeksales.sun];
          this.series=[{name:"series-1",data:salesdate}]

          const salesManSales = response.data.daily_arr;
            const salesData = salesManSales.map(sale => sale.net);
            const salesLabels = salesManSales.map(sale => sale.user_id);
            this.series1 = salesData;
            this.options1.labels = salesLabels;

            const thisMonthSales = response.data.this_month_arr;
            const salesData1 = thisMonthSales.map(sale => sale.net);
            const salesLabels1 = thisMonthSales.map(sale => sale.Userid);
            console.log("labels",salesLabels1)
            this.series2 = [{
              name: 'series-1',
              data: salesData1,
            }];
            this.options2.xaxis2.categories = salesLabels1;

    
       


      // response.data.this_month_arr.forEach((item) => {
      //   this.series2[0].data.push(item.net);
      // });

          // response.data.daily_arr.forEach((val)=>{
          //   this.series1.push(val.net)
            
          // })





          
          // this.series1=[{name: 'Sales',data: Object.values(response.data.daily_arr)}]
          // console.log("series1",this.series1)
          // this.weeklyChartData = [{ name: 'Sales', data: Object.values(weeklySalesData) }];
          // const weeklySalesData = response.data.weeklySalesArr;

          // const weeklyArr = response.data.week_arr[0];
          // this.series = [{ name: "Sales", data: Object.values(weeklyArr) }];
        
        }else{
          console.log("error")
        }
      })
      .catch((error) => {
        console.log(error);
      });
  },
};
</script>
