<template>
  <div id="app">
    <h1>My User API </h1>
    <button v-show = "!isShow" @click="isShow = !isShow, getUsers()"> Show Users </button>
<!-- //  -->
  <div v-show="isShow">
    <div class = "userTable">
      <div class = "tableId"> <b>User Id</b> </div>
      <div class = "tableFirstName"> <b>First Name</b> </div>
      <div class = "tableActions"> <b> Actions </b> </div>
    </div>
      <div class = "userTable" v-for="i in newVar" :key = i.u_id >
        <div class="tableId">
          {{i.u_id}}
        </div>
        <div class="tableFirstName">
          {{i.first_name}}
        </div>
        <div class="tableActions">
          <button> Show more </button>
                &emsp;
          <button> Edit </button>
                &emsp;
          <button v-on:click="deleteUserId(i.u_id), deleteSwitch = !deleteSwitch "> Delete User </button>
        </div>
      </div>
      <div>
        <button @click="formSwitch = !formSwitch"> New User </button>
          <!-- <div v-show="formSwitch">
            <form>
              <label for="u_id">User Id:</label><br>
              <input type="text" id="first_name" name="fname"><br>
              <label for="lname">Last name:</label><br>
              <input type="text" id="lname" name="lname">
            </form>
          </div> -->
          <div v-if="deleteSwitch">
            <!-- <div> ADDRESS </div> -->
            <p> {{reqResponse}} </p>
          </div>
      </div>
  </div>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  name: "App",
  data() {
    return {
      newVar: [],
      specificUser: [],
      isShow: Boolean,
      reqResponse: '',
      deleteSwitch: Boolean,
      formSwitch: Boolean,
    };
  },
  mounted() {
    // console.log(axios);
    // this.getUsers();
    this.function();
  },
  methods: {
    async getUsers() {
      // try {
      //   const response = await axios.get("http://localhost:5434/v1/users/");
      //   this.newVar = response.data
      //   console.log(response.data);
      // } catch(e) {
      // } finally {
      // }
      axios.get("http://localhost:5434/v1/users/").then( response =>
        this.newVar = response.data
        // console.log(this.newVar);
        // console.log(response.data),
        ).catch ( err => {
          console.log(err);
        }
        )
    },
    function(){
      this.isShow = false;
      this.switch0 = false;
      this.deleteSwitch = false;
      this.formSwitch = false;
    },
    testGetIdFunction(id) {
      var url = 'http://localhost:5434/v1/users/' + id
      axios.get(url). then (response =>
        this.specificUser = response.data
      )
    },
    async deleteUserId (id) {
      var url = 'http://localhost:5434/v1/users/' + id + '/'
      await axios.delete(url).then(
        response =>
        this.reqResponse = response.data,
      );
      this.getUsers();
      },
    },
    // async newUser (id) {
    //   var url = 'http://localhost:5434/v1/users/' + id + '/'
    //   await axios.put(url, INPUT_HERE_IN_JSON_FORMAT).then( response =>
    //   // console.log(response);
    //   this.reqResponse = response.data,
    //   );
    // }
};
</script>
<style>
/* .userTable {
  text-align: left;
  background: aqua;
} */
/* #userTable {
  text-align: right;
} */
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2C3E50;
  margin-top: 60px;
}
.userTable {
  display: flex;
  margin-bottom: 20px;
  border: 1px dashed #595959;
  padding: 8px;
  align-items: center;
  font-family: 'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
}
.tableId {
  display: flex;
  width: 5%;
}
.tableFirstName {
  display: flex;
  width: 60%;
}
.tableActions {
  display: flex;
  width: 45%;
}
</style>