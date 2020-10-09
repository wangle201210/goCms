<template>
  <div>
    <el-button class="m-10" type="primary" @click="showAdd">新增</el-button>
    <el-input
      class="w-200"
      placeholder="请输入用户名查找"
      v-model="searchName"
      clearable>
    </el-input>
    <el-table
      :data="tableData"
      border
      style="width: 100%">
      <el-table-column
        fixed
        prop="id"
        width="100"
        label="ID">
      </el-table-column>
      <el-table-column
        prop="name"
        label="姓名">
      </el-table-column>
      <el-table-column
        prop="email"
        label="邮箱">
      </el-table-column>
      <el-table-column
        prop="introduction"
        label="简介">
      </el-table-column>
      <el-table-column
        prop="created_at"
        label="创建日期">
      </el-table-column>
      <el-table-column
        fixed="right"
        width="200"
        label="操作">
        <template slot-scope="scope">
          <el-button class="m-r-5" @click="editUser(scope.row)" plain type=""primary size="small">编辑</el-button>
          <el-popconfirm
            title="确定删除？"
            @onConfirm="deleteUser(scope.row)"
          >
            <el-button slot="reference" plain type="warning" size="small">删除</el-button>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      class="f-r m-10"
      layout="prev, pager, next"
      :page-size="$config.pageSize"
      @current-change="pageChange"
      :total="total">
    </el-pagination>
    <el-dialog :visible.sync="show" title="新增管理">
      <add ref="add" @addUser="addUser" :initData="editData" :func="func"></add>
    </el-dialog>
  </div>
</template>

<script>
import {getList,addUser,editUser,deleteUser} from "@/api/user"
import add from "./add"

export default {
  name: 'example',
  components: { add },
  data() {
    return {
      tableData: [],
      show: false,
      total: 0,
      pageNum: 1,
      editData: null,
      searchName: "",
      func: 'add'
    }
  },
  created() {
  },
  mounted() {
    this.getUserList()
  },
  watch: {
    show() {
      if (!this.show) {
        this.$refs.add.resetForm()
      }
    },
    searchName() {
      this.getUserList()
    }
  },
  methods: {
    deleteUser(row) {
      deleteUser(row.id).then((res) => {
        if (res.code === 200) {
          this.getUserList()
          this.$message.success('修改成功！')
        }
      })
    },
    getUserList() {
      var params = {
        pageNum: this.pageNum
      }
      // todo 添加模糊查询
      if (this.searchName) {
        params.name = this.searchName
      }
      getList(params).then((res) => {
        if (res.code === 200) {
          this.tableData = res.data.list
          this.total = res.data.total
        }
      })
    },
    pageChange(pageNum) {
      this.pageNum = pageNum
      this.getUserList()
    },
    showAdd() {
      this.func = 'add'
      this.show = true
    },
    addUser(info) {
      if (this.func === 'add') {
        addUser(info).then((res) => {
          if (res.code === 200) {
            // 刷新用户列表
            this.getUserList()
            // 重置弹框内容
            this.$refs.add.resetForm()
            // 关闭弹框
            this.show = false
            // 提示成功
            this.$message.success('添加成功！')
          }
        })
      } else {
        editUser(this.editData.id,info).then((res) => {
          if (res.code === 200) {
            this.getUserList()
            this.$refs.add.resetForm()
            this.show = false
            this.$message.success('修改成功！')
          }
        })
      }
    },
    editUser(info) {
      this.editData = info
      this.func = 'edit'
      this.show = true
    }
  }
}
</script>

<style scoped>

</style>
