<template>
  <div>
    <input type="text" v-model="search" placeholder="Search users">
    <el-table :data="filteredUsers">
      <el-table-column prop="name" label="Name"></el-table-column>
      <el-table-column prop="email" label="Email"></el-table-column>
      <el-table-column label="Actions">
        <template #default="{ row }">
          <el-button @click="editUser(row)">Edit</el-button>
          <el-button type="danger" @click="deleteUser(row)">Delete</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page="currentPage"
      :page-sizes="[10, 20, 30, 40]"
      :page-size="pageSize"
      layout="total, sizes, prev, pager, next, jumper"
      :total="totalUsers"
    ></el-pagination>
  </div>
</template>

<script>
import { ref, computed } from 'vue';

export default {
  name: 'UserList',
  props: {
    users: Array
  },
  setup(props) {
    const currentPage = ref(1);
    const pageSize = ref(10);
    const search = ref('');
    
    const filteredUsers = computed(() => {
      return props.users.filter(user => {
        return user.name.toLowerCase().includes(search.value.toLowerCase()) ||
               user.email.toLowerCase().includes(search.value.toLowerCase());
      });
    });

    const handleSizeChange = (val) => {
      pageSize.value = val;
    };

    const handleCurrentChange = (val) => {
      currentPage.value = val;
    };

    const editUser = (user) => {
      // TODO: Implement edit user functionality
      console.log('Edit user:', user)
    };

    const deleteUser = (user) => {
      // TODO: Implement delete user functionality
      console.log('Delete user:', user)

    };

    return {
      search,
      filteredUsers,
      currentPage,
      pageSize,
      totalUsers: props.users.length,
      handleSizeChange,
      handleCurrentChange,
      editUser,
      deleteUser
    };
  }
};
</script>