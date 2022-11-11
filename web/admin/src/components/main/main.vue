<template>
  <div class="layout" :class="{'layout-hide-text': spanLeft < 5}">
    <Row type="flex">
      <i-col :span="spanLeft" class="layout-menu-left">
        <Menu :active-name="curRouter" theme="dark" width="auto">
          <div class="layout-logo-left">
            进群助手
          </div>
          <Menu-item :name="item.name" :to="item.path" v-for="(item,k) in menuList">
            <Icon type="ios-navigate" :size="iconSize"></Icon>
            <span class="layout-text">{{item.title}}</span>
          </Menu-item>
        </Menu>
      </i-col>
      <i-col :span="spanRight">
        <div class="layout-header">
          <i-button type="text" @click="toggleClick">
            <Icon type="navicon" size="32"></Icon>
          </i-button>
        </div>
        <div class="layout-breadcrumb">
          <Breadcrumb>
            <Breadcrumb-item href="#">首页</Breadcrumb-item>
            <Breadcrumb-item>{{title}}</Breadcrumb-item>
          </Breadcrumb>
        </div>
        <div class="layout-content">
              <router-view />
        </div>
        <div class="layout-copy">
          2011-2016 &copy; TalkingData
        </div>
      </i-col>
    </Row>
  </div>
</template>
<script>
export default {
  data () {
    return {
      title:"",
      spanLeft: 5,
      spanRight: 19,
      curRouter: '',
      menuList:[
        {title:"看板",path:"/dashboard",name:"dashboard"},
        {title:"群组",path:"/group",name:"group"},
        {title:"设置",path:"/setting",name:"setting"}
      ]
    }
  },
  computed: {
    iconSize () {
      return this.spanLeft === 5 ? 14 : 24;
    }
  },
  methods: {
    toggleClick () {
      if (this.spanLeft === 5) {
        this.spanLeft = 2;
        this.spanRight = 22;
      } else {
        this.spanLeft = 5;
        this.spanRight = 19;
      }
    }
  },
  mounted() {
    this.curRouter = this.$router.currentRoute.name
  },
  created() {
    this.$watch(
        () => this.$route.params,
        (toParams, previousParams) => {
          // 对路由变化做出响应...
          this.title = this.$route.meta.title
        }
    )
  },
}
</script>
<style scoped>
html,body{
  min-height:76vh;
}
.layout{
  border: 1px solid #d7dde4;
  background: #f5f7f9;
  position: relative;
  border-radius: 4px;
  height: 100%;
}
.layout-breadcrumb{
  padding: 10px 15px 0;
}
.layout-content{
  min-height: 76vh;
  margin: 15px;
  overflow: hidden;
  background: #fff;
  border-radius: 4px;
}
.layout-content-main{
  padding: 10px;
  min-height:76vh;
}
.layout-copy{
  text-align: center;
  padding: 10px 0 20px;
  color: #9ea7b4;
}
.layout-menu-left{
  background: #001529;
}
.layout-header{
  height: 60px;
  background: #fff;
  box-shadow: 0 1px 1px rgba(0,0,0,.1);
}
.layout-logo-left{
  width: 90%;
  height: 30px;
  background: #5b6270;
  border-radius: 3px;
  margin: 15px auto;
}
.layout-ceiling-main a{
  color: #9ba7b5;
}
.layout-hide-text .layout-text{
  display: none;
}
.ivu-col{
  transition: width .2s ease-in-out;
}
.layout-text a {
  color: #fff;
}
.layout-text  a.router-link-active {
  color: #2D8cF0;
}
</style>
