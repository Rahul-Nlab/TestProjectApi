<template>

  <div id="app">

    <h1>My User API </h1>
    <button v-show = "!isShow" @click="isShow = !isShow, getUsers()"> Show Users </button>

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
              <button v-on:click="deleteUserId(i.u_id), setShow()"> Delete User </button>

            </div>

          </div>

          <div >

            <button @click="newUser(id)"> New User </button>
    <!-- HERE I HAVE TO CREATE USER DETAIL FORM THAT POPS UP ONLY WHEN CREATE USER IS CLICKED, AND DISPLAYS A FORM THAT ACCEPTS DATA AND CALLS A FUNCTION ON SUBMIT CLICK WITH A JSON FORMAT PARAMETER -->
    <!-- ADD A SWITCH THAT SHOWS EITHER THE FORM OR THE NEW USER BUTTON -->
            <div class = "formStyle">

                <form>
                  Enter Details: <br>

                  <input v-model="userId" type="text" id="id" name="id" placeholder='User Id'><br>
                  <input v-model="firstName" type="text" id="fname" name="fname" placeholder='First Name'><br>
                  <input v-model="middleName" type="text" id="mname" name="mname" placeholder='Middle Name'><br>
                  <input v-model="lastName" type="text" id="lname" name="lname" placeholder='Last name'>

                  <button @click="newUser()"> Create </button>

                </form>

            </div>

          </div>

          <!-- <div> ADDRESS

          </div> -->

    <!-- THIS IS DEDICATED DELETE MESSAGE SHOWING DIV -->
          <div v-show="isShowDelete">

            <p> {{reqResponse}} </p>

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
      reqResponse: '',
      isShowDelete: Boolean,
      formSwitch: Boolean,
      isShow: Boolean,
      userId: "",
      firstName: "",
      middleName: "",
      lastName: "",
    };
  },

  mounted() {
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

    setShow () {
      this.isShowDelete = true;
      setTimeout(() => {
        this.isShowDelete = false;
      }, 1500);
    },

    async newUser () {
      // console.log(this.firstName, this.middleName, this.lastName);
      let url
      if (this.userId === '' ) {
        url = 'http://localhost:5434/v1/users/'
      } else {
        url = 'http://localhost:5434/v1/users/' + this.userId + '/'
      }

      // let jsonRes = {"first_name": this.firstName, "middle_name": this.middleName, "last_name": this.lastName};

      let jsonRes = '{ "first_name":' + '"' + this.firstName + '",' + '"middle_name":' + '"' + this.middleName + '",' +  '"last_name":' + '"' + this.lastName + '" };'

      console.log(jsonRes);
      await axios.post(url, jsonRes).then( response =>
      this.reqResponse = response.data,
      // console.log(response.data),
      );
    },
  },
};

</script>


<style>

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

.formStyle {
  display: flex;
  width: 300px;
  font-size: medium;
  margin-bottom: 20px;
  border: 1px dashed #595959;
  padding: 8px;
  align-items: left;
  font-family: 'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
}

</style>