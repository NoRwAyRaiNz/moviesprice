<template>
  <div>
    <el-menu
      :default-active="activeIndex"
      class="el-menu-demo"
      mode="horizontal"
      @select="handleSelect"
      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b"
    >
      <el-menu-item index="1">首页</el-menu-item>
      <el-submenu index="2">
        <template slot="title">地区</template>
        <el-menu-item index="2-1">中国</el-menu-item>
        <el-menu-item index="2-2">美国</el-menu-item>
        <el-menu-item index="2-3">日本</el-menu-item>
        <el-submenu index="2-4">
          <template slot="title">欧洲</template>
          <el-menu-item index="2-4-1">英国</el-menu-item>
          <el-menu-item index="2-4-2">法国</el-menu-item>
          <el-menu-item index="2-4-3">德国</el-menu-item>
        </el-submenu>
      </el-submenu>
    </el-menu>
  </div>
</template>

<script>
export default {
  data () {
    return {
      activeIndex: '1'
    }
  },
  methods: {
    handleSelect (key, keyPath) {
      console.log(key, keyPath)
    }
  }
}
</script>

<style>
</style>
