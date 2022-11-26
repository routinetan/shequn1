<template>
  <div class="wrapper">
    <router-view />
  </div>
</template>
<script>
export default {};
</script>

<style lang="less">
html,body{
  height:100%;
  width:100%;
  font-size:24px;
}
#app{
  padding:0px;
  margin:0px;
  position:absolute;
  top:0px;
  left:0px;
  width:100%;
  height:100%;
  border:hidden;
}
.wrapper {
  height:100%;
}
</style>
