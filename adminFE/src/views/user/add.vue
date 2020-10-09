<template>
  <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
    <el-form-item label="用户名" prop="name">
      <el-input v-model="ruleForm.name"></el-input>
    </el-form-item>
    <el-form-item label="邮箱" prop="email">
      <el-input v-model="ruleForm.email"></el-input>
    </el-form-item>
    <el-form-item label="密码" prop="password">
      <el-input :placeholder="passwordHolder" v-model="ruleForm.password"></el-input>
    </el-form-item>
    <el-form-item label="头像" prop="avatar">
      <avatar :image="ruleForm.avatar" @success="avatarUp"/>
    </el-form-item>
    <el-form-item label="角色" prop="role">
      <el-select v-model="ruleForm.role" placeholder="请选择角色">
        <el-option label="超级管理员" :value="1"></el-option>
        <el-option label="内容管理员" :value="2"></el-option>
      </el-select>
    </el-form-item>
    <el-form-item label="简介" prop="introduction">
      <el-input type="textarea" v-model="ruleForm.introduction"></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="submitForm('ruleForm')">保存</el-button>
      <el-button @click="resetForm('ruleForm')">重置</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import avatar from "@/components/Upload/avatar"

export default {
  components: { avatar },
  name: 'userAdd',
  props:{
    initData: {
      type: Object,
      default: () => {}
    },
    func: {
      type: String,
      default: 'add'
    }
  },
  data() {
    return {
      ruleForm: {
        name: '',
        email: '',
        password: '',
        role: '',
        avatar: '',
        introduction: '',
      },
      rules: {
        name: [
          { required: true, message: '请输入活动名称', trigger: 'blur' },
          { min: 3, max: 10, message: '长度在 3 到 10 个字符', trigger: 'blur' }
        ],
        // todo 添加邮箱格式验证
        email: [
          { required: true, message: '邮箱必填', trigger: 'blur' },
        ],
        role: [
          { required: true, message: '请选择角色', trigger: 'change' }
        ],
        password: [
          { min: 5, max: 10, message: '长度在 5 到 10 个字符', trigger: 'blur' }
        ]
      }
    };
  },
  computed:{
    passwordHolder() {
      if (this.func === 'add') {
        return '不填则为系统默认密码password'
      } else {
        return '不填写则不修改'
      }
    }
  },
  methods: {
    submitForm() {
      this.$refs.ruleForm.validate((valid) => {
        if (valid) {
          this.$emit("addUser",this.ruleForm)
        } else {
          console.log('error submit!!');
          return false;
        }
      });
    },
    resetForm() {
      this.$refs.ruleForm.resetFields();
    },
    copyData() {
      if (this.initData) {
        this.ruleForm = _.pick(this.initData,['name','email','avatar','introduction','role','password'])
        this.ruleForm.password = ''
      }
    },
    avatarUp(url) {
      this.ruleForm.avatar = url
    }
  },
  mounted() {
    this.copyData()
  },
  watch: {
    initData: {
      handler() {
        this.copyData()
      },
      deep:true
    }
  }
}
</script>
<style scoped>

</style>
