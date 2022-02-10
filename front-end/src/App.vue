<template>

  <div id="app">

    <h1>My User API </h1>
    <button v-show = "!isShow" @click="isShow = !isShow, getUsers()"> Show Users </button>

      <div v-show="isShow">

        <div class = "userTable">

          <div class = "tableId"> <b>User Id</b> </div>

          <div class = "tableFirstName"> <b>First Name</b> </div>

          <div class = "tableMiddleName"> <b>Middle Name</b> </div>

          <div class = "tableLastName"> <b>Last Name</b> </div>

          <div class = "tableActions"> <b> Actions </b> </div>

        </div>

        <div class = "userTable" v-for="i in newVar" :key = i.u_id >

          <div class="tableId">

            {{i.u_id}}

          </div>

          <div class="tableFirstName">

            {{i.first_name}}

          </div>


          <div class="tableMiddleName">

            {{i.middle_name}}

          </div>

          <div class="tableLastName">

            {{i.last_name}}

          </div>

          <div class="tableActions">

            <button> Show Address </button>
                  &emsp;
            <button v-on:click="setEditValues(i), showEditForm()"> Edit </button>
            <!-- editUserId(i.u_id), -->
                  &emsp;
            <button v-on:click="deleteUserId(i.u_id), setShow()"> Delete User </button>

          </div>

        </div>

        <div >

          <button @click="formSwitch = !formSwitch" v-show="formSwitch"> New User </button>

  <!-- HERE I HAVE TO CREATE USER DETAIL FORM THAT POPS UP ONLY WHEN CREATE USER IS CLICKED, AND DISPLAYS A FORM THAT ACCEPTS DATA AND CALLS A FUNCTION ON SUBMIT CLICK WITH A JSON FORMAT PARAMETER -->
  <!-- ADD A SWITCH THAT SHOWS EITHER THE FORM OR THE NEW USER BUTTON -->
          <div class = "formStyle" v-show="!formSwitch">

              <form>
                Enter Details: <br>

                <input v-model="userId" type="text" id="id" name="id" placeholder='User Id'><br>
                <input v-model="firstName" type="text" id="fname" name="fname" placeholder='First Name'><br>
                <input v-model="middleName" type="text" id="mname" name="mname" placeholder='Middle Name'><br>
                <input v-model="lastName" type="text" id="lname" name="lname" placeholder='Last name'><br>

                <button @click="newUser(), preventEvent($event), setShow(), formSwitch = !formSwitch"> Create </button>
                &emsp;
                <button @click="formSwitch = !formSwitch, preventEvent($event)"> Cancel </button>

              </form>

          </div>

        </div>


      </div>

<!-- formSwitchForEdit IS YET TO BE DECLARED, IT IS JUST ANOTHER VARIABLE FOR MAKING EDIT FORM VISIBLE AND INVISIBLE. placeholder='First Name' -->

        <div class = "formStyle" v-show="formSwitchForEdit"> 
          <!-- A DIV FOR EDIT FORM -->
          
          <form> 
            Edit details: <br> 

            <!-- <input v-model="firstName" type="text" id="fname" name="fname"  value=newVarFirstName><br> -->
            <input v-model="firstNameForEdit" type="text" id="fname" name="fname"><br>
            <input v-model="middleNameForEdit" type="text" id="mname" name="mname"><br>
            <input v-model="lastNameForEdit" type="text" id="lname" name="lname"><br>

            <button @click="closeEditForm(), editUserId(), preventEvent($event), setShow()"> Update </button>
            &emsp;
            <button @click="preventEvent($event), closeEditForm()"> Cancel </button>

          </form>

        </div>
          <!-- <div> ADDRESS

          </div> -->

    <!-- THIS IS DEDICATED DELETE MESSAGE SHOWING DIV -->
          <div v-show="isShowDelete">

            <p> {{reqResponse}} </p>

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
      // newVar1: {},
      newVarFirstName: '',
      specificUser: [],
      reqResponse: '',
      userIdForGet: '',
      formSwitchForEdit: Boolean,
      isShowDelete: Boolean,
      formSwitch: Boolean,
      isShow: Boolean,
      userId: "",
      firstName: "",
      middleName: "",
      lastName: "",
      userIdForEdit: "",
      firstNameForEdit: "",
      middleNameForEdit: "",
      lastNameForEdit: "",
    };
  },

  mounted() {
    // this.getUsers();
    this.function();
  },
  
    // try {
    //   const response = await axios.get("http://localhost:5434/v1/users/");
    //   this.newVar = response.data
    //   console.log(response.data);
    // } catch(e) {
    // } finally {
    // }
    // v-show = "formSwitchForEdit"

  methods: {
    
    
    function(){
      this.isShow = false;
      this.switch0 = false;
      this.formSwitch = true;
      this.deleteSwitch = false;
      this.formSwitchForEdit = false;
    },

    preventEvent(event) {
      event.preventDefault();
    },

    showEditForm() {
      this.formSwitchForEdit = true;
    },
    closeEditForm() {
      this.formSwitchForEdit = false;
    },

    // testGetIdFunction(id) {
    //   var url = 'http://localhost:5434/v1/users/' + id
    //   axios.get(url). then (response =>
    //     this.specificUser = response.data
    //   )
    // },

    async deleteUserId (id) {
      var url = 'http://localhost:5434/v1/users/' + id + '/'
      await axios.delete(url).then(
        response =>
        this.reqResponse = response.data,
      );
      this.getUsers();
    },

    async editUserId() {

      var url = 'http://localhost:5434/v1/users/' + this.userIdForEdit + '/'
      let jsonRes = '{ "first_name":' + '"' + this.firstNameForEdit + '",' + '"middle_name":' + '"' + this.middleNameForEdit + '",' +  '"last_name":' + '"' + this.lastNameForEdit + '" }'
// console.log(jsonRes);
      await axios.put(url,jsonRes).then(
          response =>
        this.reqResponse = response.data,
      );
      this.getUsers();
    },

    setEditValues(i) {
      this.userIdForEdit = i.u_id;
      this.firstNameForEdit = i.first_name;
      this.middleNameForEdit = i.middle_name;
      this.lastNameForEdit = i.last_name;
    },

    async newUser () {

      let url
      
      if (this.userId === '' ) {
        url = 'http://localhost:5434/v1/users/'
      } else {
        url = 'http://localhost:5434/v1/users/' + this.userId + '/'
      }

      let jsonRes = '{ "first_name":' + '"' + this.firstName + '",' + '"middle_name":' + '"' + this.middleName + '",' +  '"last_name":' + '"' + this.lastName + '" }'

      await axios.post(url, jsonRes).then( response =>
      this.reqResponse = response.data,
      );

      this.userId = '';
      this.firstName = '';
      this.middleName = '';
      this.lastName = '';

      this.getUsers();

    },

    async getUsers() {
      await axios.get('http://localhost:5434/v1/users/').then( response =>
        this.newVar = response.data
          ).catch ( err => {
            console.log(err);
          }
        )
    },

    setShow () {
      this.isShowDelete = true;
      setTimeout(() => {
        this.isShowDelete = false;
      }, 1500);
    },


  },
};

</script>


<style>

#app {
  font-family: "Courier New", Courier, monospace;
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
}

.tableId {
  display: flex;
  width: 5%;
}

.tableFirstName {
  display: flex;
  width: 10%;
}

.tableMiddleName {
  display: flex;
  width: 10%;
}

.tableLastName {
  display: flex;
  width: 50%;
}

.tableActions {
  display: flex;
  width: 25%;
}

.formStyle {
  display: flex;
  width: 10%;
  font-size: medium;
  margin-bottom: 20px;
  border: 1px dashed #595959;
  padding: 10px;
  line-height: 25px;
}

button{
  font-family: "Courier New", Courier, monospace;
}

input{
    font-family: "Courier New", Courier, monospace;
}

::placeholder{
  font-family: "Courier New", Courier, monospace;
}

</style>