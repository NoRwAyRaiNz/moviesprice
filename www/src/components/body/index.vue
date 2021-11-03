<template>
  <div class="main">
    <el-tabs v-model="activeName" @tab-click="handleClick">
    <el-tab-pane label="全部" name="first">
        <div class="content">
            <img src="@/assets/post.jpg">
            <div class="price">$19.99</div>
            <div class="name">测试测试</div>
        </div>
    </el-tab-pane>
    <el-tab-pane label="推荐" name="second">推荐</el-tab-pane>
    <el-tab-pane label="热门" name="third">热门</el-tab-pane>
    <el-tab-pane label="即将上映" name="fourth">即将上映</el-tab-pane>
  </el-tabs>
  </div>
</template>

<script>
export default {
  data () {
    return {
      activeName: 'second'
    }
  },
  methods: {
    handleClick (tab, event) {
      console.log(tab, event)
    }
  }
}
</script>

<style scoped>
.main{
    margin:0 auto;
    margin-top: 2%;
    width: 80%;
}
.content{
    width: 15%;
    padding: 10px;
    border: 1px solid #ccc;
}
.content img{
    width: 184px;
    height: 277px;
}
.content .price{
    margin-top: 10px;
    font-size: 25px;
    color: #2196f3;
    text-align: center;
}
.content .name{
    margin-top: 10px;
    text-align: center;
}
</style>
