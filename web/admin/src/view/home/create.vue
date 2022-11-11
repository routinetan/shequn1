<template>
<Row>
  <Row>
    <Menu mode="horizontal"  active-name="1">
      <MenuItem name="1" to="/group/1/edit_nav1">
        基础
      </MenuItem>
      <MenuItem name="2" to="/group/1/edit_nav2">
        高级
      </MenuItem>
      <MenuItem name="4" to="/group/1/edit_nav3">
        分销
      </MenuItem>
    </Menu>
  </Row>
  <Row>
    <Col span="12" class="info_content">
      <router-view />
    </Col>
  </Row>
</Row>
</template>
<script>
export default {
  data () {
    return {

    }
  }
}
</script>
<style scoped>
.info_content {
  padding: 20px 0;
}
</style>
