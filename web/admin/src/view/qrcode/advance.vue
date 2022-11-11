<template>
  <Row>
    <Col span="12">


      <FormItem label="是否收费">
        <i-switch v-model="formItem.hot_status" size="large" :true-value="trueValue" :false-value="falseValue" >
          <span slot="open">是</span>
          <span slot="close">否</span>
        </i-switch>
      </FormItem>
      <FormItem label="价格">
        <Input v-model="formItem.title" placeholder="请输入6个以内的汉字"></Input>
      </FormItem>
      <FormItem label="群最大容量">
        <Input v-model="formItem.title" placeholder="请输入6个以内的汉字"></Input>
      </FormItem>
      <FormItem label="子群数">
        <Input v-model="formItem.title" placeholder="请输入6个以内的汉字"></Input>
      </FormItem>
      <FormItem label="群扫码次数">
        <Input v-model="formItem.title" placeholder="请输入6个以内的汉字"></Input>
      </FormItem>
      <FormItem label="群有效期">
        <Input v-model="formItem.title" placeholder="请输入6个以内的汉字"></Input>
      </FormItem>
      <FormItem label="佣金比率">
        <Input v-model="formItem.title" placeholder="请输入6个以内的汉字"></Input>
      </FormItem>
      <FormItem label="群类型">
        <i-switch v-model="formItem.hot_status" size="large" :true-value="trueValue" :false-value="falseValue" >
          <span slot="open">是</span>
          <span slot="close">否</span>
        </i-switch>
      </FormItem>
      <FormItem label="是否置顶">
        <i-switch v-model="formItem.top_status" size="large" :true-value="trueValue" :false-value="falseValue" >
          <span slot="open">是</span>
          <span slot="close">否</span>
        </i-switch>
      </FormItem>
      <FormItem label="是否设置为热门群">
        <i-switch v-model="formItem.hot_status" size="large" :true-value="trueValue" :false-value="falseValue" >
          <span slot="open">是</span>
          <span slot="close">否</span>
        </i-switch>
      </FormItem>
    </Col >
  </Row >
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
        title: "第6群",
        max_num: 200,
        base_num:198,
        status:1,
        qrcode_url:''
      }
    }
  },
  methods: {
    handleIconSuccess(response, file, fileList, index) {
      this.thumb_url = response.data.prefix + response.data.file_list.file[0]
      this.formItem.qrcode_url = response.data.file_list.file[0]
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
        url:'/qrcode/add',
        data:this.formItem,
        method:"POST",
      }).then(response => {
        this.$Modal.info({
          content:"提交成功",
          onOk:()=>{
            this.$router.push("/group/2/qrcode")
          }
        })
      })
    }
  },
  created() {
    this.group_id = parseInt(this.$route.params.id)
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
