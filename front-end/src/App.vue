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

      <!-- <div> -->

        <div v-for="i in newVar" :key = i.u_id >
          <div class = "userTable">

            <div class="tableId"> {{i.u_id}} </div>
            <div class="tableFirstName"> {{i.first_name}} </div>
            <div class="tableMiddleName"> {{i.middle_name}} </div>
            <div class="tableLastName"> {{i.last_name}} </div>

            <div class="tableActions">

              <button @click="getUserAddresses(i), showAddForm(), closeEditForm(), setShow(), hideNewForm(), j=i"> Show Address </button>
                    &emsp;
              <button v-on:click="setEditValues(i), showEditForm(), hideAddForm(), hideNewForm()"> Edit </button>
                    &emsp;
              <button v-on:click="deleteUserId(i.u_id), setShow(), hideAddForm(), closeEditForm(), hideNewForm()"> Delete User </button>
              
            </div>

          </div>

          <div v-if="isShowAddresses & i===j" class = "addressTable" >
            <div v-for="j in newVarAddresses" :key = j.a_id class = "addressSpecific" >
                
                <br>Street: {{j.street}}
                <br>Area: {{j.area}} 
                <br>Pincode: {{j.pincode}} 
                <br>City: {{j.city}}
                <!-- &emsp; -->
              
            </div>
          </div>

        </div>

       <!-- </div>  -->
          
        <div>
          <button @click="showNewForm(), hideAddForm()" v-show="!switchForNew"> New User </button>
        </div>

  <!-- HERE I HAVE TO CREATE USER DETAIL FORM THAT POPS UP ONLY WHEN CREATE USER IS CLICKED, AND DISPLAYS A FORM THAT ACCEPTS DATA AND CALLS A FUNCTION ON SUBMIT CLICK WITH A JSON FORMAT PARAMETER -->
  <!-- ADD A SWITCH THAT SHOWS EITHER THE FORM OR THE NEW USER BUTTON -->
          <div class = "formStyle" v-show="switchForNew">

            <form>
              Enter Details: <br>

              <input v-model="userId" type="text" id="id" name="id" placeholder='User Id'><br>
              <input v-model="firstName" type="text" id="fname" name="fname" placeholder='First Name'><br>
              <input v-model="middleName" type="text" id="mname" name="mname" placeholder='Middle Name'><br>
              <input v-model="lastName" type="text" id="lname" name="lname" placeholder='Last name'><br>

              <button @click="newUser(), preventEvent($event), setShow(), hideNewForm()"> Create </button>
              &emsp;
              <button @click="hideNewForm(), preventEvent($event)"> Cancel </button>

            </form>

          </div>

        </div>

      <!-- </div> -->

      <div class = "formStyle" v-show="SwitchForEdit"> 
        <!-- a DIV FOR EDIT FORM -->     

        <form> 
          Edit details: <br> 

          <!-- <input v-model="firstName" type="text" id="fname" name="fname"  value=newVarFirstName><br> -->
          <input v-model="firstNameForEdit" type="text" id="fname" name="fname"><br>
          <input v-model="middleNameForEdit" type="text" id="mname" name="mname"><br>
          <input v-model="lastNameForEdit" type="text" id="lname" name="lname"><br>

          <button @click="closeEditForm(), preventEvent($event), editUserId(), setShow()"> Update </button>
                &emsp;
          <button @click="preventEvent($event), closeEditForm()"> Cancel </button>

        </form>

      </div>

      <!-- THIS IS DEDICATED DELETE MESSAGE SHOWING DIV -->
      <div v-show="isShowResponse">   <p> {{reqResponse}} </p>    </div>  

  </div>

</template>
// S------------------------------C------------------------------R------------------------------I------------------------------P------------------------------T------------------------------           
<script>

import axios from 'axios';

export default {
  name: "App",

  data() {
    return {
      newVar: [],
      specificUser: [],
      newVarAddresses: [],
      reqResponse: '',
      userIdForGet: '',
      newVarFirstName: '',
      isShow: Boolean,
      switchForNew: Boolean,
      SwitchForEdit: Boolean,
      isShowResponse: Boolean,
      isShowAddresses: Boolean,
      userId: "",
      lastName: "",
      firstName: "",
      middleName: "",
      userIdForEdit: "",
      lastNameForEdit: "",
      firstNameForEdit: "",
      middleNameForEdit: "",
      increment: Number,
    };
  },

  mounted() {
    this.function();
  
    // try {
      //   const response = await axios.get("http://localhost:5434/v1/users/");
    //   this.newVar = response.data
    //   console.log(response.data);
    // } catch(e) {
      // } finally {
        // }
    // v-show = "SwitchForEdit"

    },

  methods: {
    
    function(){
      this.isShow = false;
      this.switch0 = false;
      this.switchForNew = false;
      this.deleteSwitch = false;
      this.SwitchForEdit = false;
      this.isShowAddresses = false;
    },

    preventEvent(event) {
      event.preventDefault();
    },

    showNewForm() {
      this.switchForNew = true;
    },

    hideNewForm() {
      this.switchForNew = false;
    },

    showEditForm() {
      this.SwitchForEdit = true;
    },

    closeEditForm() {
      this.SwitchForEdit = false;
    },

    showAddForm() {
      this.isShowAddresses = true;
    },

    hideAddForm() {
      this.isShowAddresses = false;
    },

    async deleteUserId (id) {
      var url = 'http://localhost:5434/v1/users/' + id + '/'
      await axios.delete(url).then(
        response =>
        this.reqResponse = response.data,
      );
      this.getUsers();
    },

    async editUserId() {
      if (this.validateNames(this.firstNameForEdit, this.middleNameForEdit, this.lastNameForEdit)) {
        var url = 'http://localhost:5434/v1/users/' + this.userIdForEdit + '/'
        let jsonRes = '{ "first_name":' + '"' + this.firstNameForEdit + '",' + '"middle_name":' + '"' + this.middleNameForEdit + '",' +  '"last_name":' + '"' + this.lastNameForEdit + '" }'
        await axios.put(url,jsonRes).then(
            response =>
          this.reqResponse = response.data,
        );
        this.getUsers();
      }
    },

    validateNames(first, middle, last) {
      // validation here
      let x = /[^A-Za-z]/
      if ( x.test(first) || x.test(middle) || x.test(last)){
        this.reqResponse = "Names cannot contain Symbols/Numbers";
        return false;

      } else if ( first == "" || middle == "" || last == ""){
        this.reqResponse = "Names cannot be empty.";
        return false;

      } else {
        // this.editUserId()
        return true;
      }

    },

    setEditValues(i) {
      this.userIdForEdit = i.u_id;
      this.firstNameForEdit = i.first_name;
      this.middleNameForEdit = i.middle_name;
      this.lastNameForEdit = i.last_name;
    },

    async newUser () {
      if (this.validateNames(this.firstName, this.middleName, this.lastName)) {
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
      }
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
      this.isShowResponse = true;
      setTimeout(() => {
        this.isShowResponse = false;
        this.reqResponse = '';
      }, 1500);
    },

    async getUserAddresses(i) {
      let url = "http://localhost:5434/v1/users/" + i.u_id + "/addresses/"
      await axios.get(url).then( response =>
        this.newVarAddresses = response.data
        ).catch( err => {
          console.log(err);
      })
      console.log(this.newVarAddresses);
      if (this.newVarAddresses.length === 0 ) {
        this.reqResponse = "No Addresses";
        console.log("true");
      }
    }, 


  },
};

</script>
// S------------------------------T------------------------------Y------------------------------L------------------------------E------------------------------
<style>

#app {
  font-family: "Courier New", Courier, monospace;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2C3E50;
  margin-top: 60px;
}

.addressTable {
  display: flex;
}

.addressSpecific {
  margin-bottom: 20px;
  border: 1px dashed #595959;
  width: fit-content;
  padding: 5px;
  text-align: left;
  display: flex;
  width: fit-content;
  margin-right: 20px;
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