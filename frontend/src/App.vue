<template>
  <div id="app">
    <UserList :users="users" />
    <UserForm :user="selectedUser" />
  </div>
</template>

<script>
import UserList from './components/UserList.vue';
import UserForm from './components/UserForm.vue';

export default {
  name: 'App',
  components: {
    UserList,
    UserForm
  },
  data() {
    return {
      users: [
        { id: 1, name: 'John Doe', email: 'john@example.com' },
        { id: 2, name: 'Jane Doe', email: 'jane@example.com' }
      ],
      selectedUser: null
    };
  }
};
</script>
