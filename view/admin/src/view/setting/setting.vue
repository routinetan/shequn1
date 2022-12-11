<template>
  <Row >
      <Card :bordered="false">
        <template #title>基础设置</template>
        <Form :model="formItem" :label-width="80">
          <FormItem label="系统名称">
            <Input v-model="formItem.title" placeholder="请输入6个以内的汉字"></Input>
          </FormItem>
          <FormItem label="是否开启VIP模式">
            <i-switch v-model="formItem.vip_status" size="large" :true-value="trueValue" :false-value="falseValue" >
              <span slot="open">开启</span>
              <span slot="close">关闭</span>
            </i-switch>
          </FormItem>
          <FormItem label="是否开启支付">
            <i-switch v-model="formItem.pay_close_status" size="large" :true-value="trueValue" :false-value="falseValue" >
              <span slot="open">开启</span>
              <span slot="close">关闭</span>
            </i-switch>
          </FormItem>
          <FormItem label="是否开始默认个微">
            <i-switch v-model="formItem.default_wechat_status" size="large" :true-value="trueValue" :false-value="falseValue" >
              <span slot="open">开启</span>
              <span slot="close">关闭</span>
            </i-switch>
          </FormItem>
          <FormItem label="默认个微二维码" v-if="formItem.default_wechat_status == 1">
            <div class="demo-upload-list" v-if="formItem.default_wechat_qrcodeurl !=''" >
              <template >
                <img :src="thumb_url">
              </template>
            </div>
            <Upload
                ref="app_upload"
                :show-upload-list="false"
                :on-success="handleWechatQrcodeSuccess"
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
          <FormItem label="是否开始炮灰域名">
            <i-switch v-model="formItem.feeder_domain_status" size="large" :true-value="trueValue" :false-value="falseValue" >
              <span slot="open">开启</span>
              <span slot="close">关闭</span>
            </i-switch>
          </FormItem>
          <FormItem label="是否开启兜底群">
            <i-switch v-model="formItem.gen_group_status" size="large" :true-value="trueValue" :false-value="falseValue" >
              <span slot="open">开启</span>
              <span slot="close">关闭</span>
            </i-switch>
          </FormItem>
          <FormItem label="默认兜底群id" v-if="formItem.gen_group_status == 1">
            <Input v-model="formItem.gen_group_id" placeholder="请输入6个以内的汉字"></Input>
          </FormItem>
          <FormItem label="是否开启系统">
            <i-switch v-model="formItem.close_status" size="large" :true-value="trueValue" :false-value="falseValue" >
              <span slot="open">开启</span>
              <span slot="close">关闭</span>
            </i-switch>
          </FormItem>
          <FormItem>
            <Button type="primary"  @click="handleSubmit()">提交</Button>
          </FormItem>
        </Form>
      </Card>
  </Row>

</template>
<script>
import config from '@/api/config.js'
export default {
  data () {
    return {
      groupId:"",
      upload:config.baseURL+"/upload",
      trueValue:1,
      falseValue:0,
      thumb_url:'',
      formItem: {
        title: "",
        gen_group_id:14,
        close_status:0,
        pay_close_status:1,
        vip_status:0,
        gen_group_status:0,
        default_wechat_status:1,
        feeder_domain_status:0,
        default_wechat_qrcodeurl:'',
      }
    }
  },
  methods: {
    handleWechatQrcodeSuccess(response, file, fileList, index) {
      this.thumb_url = response.data.prefix + response.data.file_list.file[0]
      this.formItem.default_wechat_qrcodeurl = response.data.file_list.file[0]
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
    handleSubmit() {
      this.formItem["group_id"] = this.group_id
      this.Request({
        url:'/systemCfg',
        data:this.formItem,
        method:"POST",
      }).then(response => {
        this.$Modal.info({
          content:"提交成功"
        })
      })
    },
  },
  mounted() {
    this.group_id = parseInt(this.$route.params.id)
    this.Request({
      url:'/systemCfg',
      method:"GET",
    }).then(response => {
      this.formItem = response.data.data
      this.thumb_url = response.data.data.thumb_url
    })
  }
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
