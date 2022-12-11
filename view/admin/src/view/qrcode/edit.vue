<template>
  <Form :model="formItem" :label-width="80">
    <FormItem label="群名">
      <Input v-model="formItem.title" placeholder="Enter something..."></Input>
    </FormItem>
    <FormItem label="当前人数(前台显示)" >
      <Input v-model="formItem.baseNum" placeholder="Enter something..."></Input>
    </FormItem>
    <FormItem label="最大人数">
      <Input v-model="formItem.maxNum" placeholder="Enter something..."></Input>
    </FormItem>
    <FormItem label="状态">
      <i-switch v-model="formItem.switch" size="large">
        <span slot="open">开放</span>
        <span slot="close">关闭</span>
      </i-switch>
    </FormItem>
    <FormItem label="活码" >
      <div class="demo-upload-list" v-if="formItem.thumbUrl !=''" >
        <template >
          <img :src="formItem.thumbUrl">
        </template>
      </div>
      <Upload
          ref="app_upload"
          :show-upload-list="false"
          :on-success="handleIconSuccess"
          :format="['jpg','jpeg','png']"
          :on-format-error="handleFormatError"
          type="drag"
          :data="{'key':'file'}"
          :action="upload"
          style="display: inline-block;width:58px;">
        <div style="width: 58px;height:58px;line-height: 58px;">
          <Icon type="md-add" size="20"></Icon>
        </div>
      </Upload>
    </FormItem>
    <FormItem>
      <Button type="primary">Submit</Button>
      <Button style="margin-left: 8px">Cancel</Button>
    </FormItem>
  </Form>
</template>
<script>
import config from '@/api/config.js'
export default {
  data () {
    return {
      groupId:"",
      upload:config.baseURL+"/upload",
      formItem: {
        title: undefined,
        maxNum: undefined,
        baseNum:undefined,
        thumbUrl:undefined
      }
    }
  },
  methods: {
    handleIconSuccess(response, file, fileList, index) {
      file.thumbUrl = response.data.file[0]
      this.formItem.icon = file.thumbUrl
    },
    handleFormatError(file) {
      this.$Notice.warning({
        title: 'The file format is incorrect',
        desc: 'File format of ' + file.name + ' is incorrect, please select jpg or png.'
      });
    },
    handleMaxSize(file) {
      this.$Notice.warning({
        title: 'Exceeding file size limit',
        desc: 'File  ' + file.name + ' is too large, no more than 2M.'
      });
    },
    handlerSubmit() {
      this.formItem["group_id"] = this.group_id
      this.Request({
        url:'qrcode/add',
        data:this.formItem,
        method:"POST",
      }).then(response => {

      })
    }
  },
}
</script>
<style scoped>
.qrcode {
  width: 103px;
  height: 103px;
}
.icon_list{
  display: inline-block;
  width: 103px;
  height: 175px;
  text-align: center;
  line-height: 60px;
  border: 1px solid transparent;
  border-radius: 4px;
  overflow: hidden;
  background: #fff;
  position: relative;
  box-shadow: 0 1px 1px rgba(0,0,0,.2);
  margin-right: 4px;
}
.icon_list img{
  width: 100%;
  height: 100%;
}
.demo-upload-list{
  display: inline-block;
  width: 120px;
  height: 120px;
  text-align: center;
  line-height: 60px;
  border: 1px solid transparent;
  border-radius: 4px;
  overflow: hidden;
  background: #fff;
  position: relative;
  box-shadow: 0 1px 1px rgba(0,0,0,.2);
  margin-right: 4px;
}
.demo-upload-list img{
  width: 100%;
  height: 100%;
}
.demo-upload-list-cover{
  display: none;
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  background: rgba(0,0,0,.6);
}
.demo-upload-list:hover .demo-upload-list-cover{
  display: block;
}
.demo-upload-list-cover i{
  color: #fff;
  font-size: 20px;
  cursor: pointer;
  margin: 0 2px;
}
</style>
