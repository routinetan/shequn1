<style scoped>
html,body{
  height:100vh;
}
.layout{
  border: 1px solid #d7dde4;
  background: #f5f7f9;
  position: relative;
  border-radius: 4px;
  overflow: hidden;
  height: 100%;
}
.layout-breadcrumb{
  padding: 10px 15px 0;
}
.layout-content{
  min-height: 200px;
  margin: 15px;
  overflow: hidden;
  background: #fff;
  border-radius: 4px;
}
.layout-content-main{
  padding: 10px;
  height:80vh;
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
</style>
<template>
    <div>
        <Row :gutter="4" >
                <Button type="primary"  to="/group/2/qrcode/add">新增活码</Button>
        </Row>
        <Br />
        <Row>
              <Table border :columns="columns1" :data="list">
              </Table>
        </Row>
    </div>
</template>
<script>
export default {
  data () {
    return {
      spanLeft: 5,
      spanRight: 19,
      title:"",
      group_id:"",
      columns1: [
        {
          title: 'id',
          key: 'id',
        },
        {
          title: '群名',
          key: 'title'
        },
        {
          title: '进群人数',
          key: 'num'
        },
        {
          title: '当前群人数',
          key: 'status'
        },
        {
          title: '状态',
          key: 'status'
        },
        {
          title: '操作',
          key: 'action',
          width: 150,
          align: 'center',
          render: (h, params) => {
            return h('div', [
              h('Button', {
                props: {
                  type: 'primary',
                  size: 'small'
                },
                style: {
                  marginRight: '5px'
                },
                on: {
                  click: () => {
                    this.show(params.index);
                  }
                }
              }, '查看'),
              h('Button', {
                props: {
                  type: 'error',
                  size: 'small'
                },
                on: {
                  click: () => {
                    this.remove(params.index)
                  }
                }
              }, '删除')
            ]);
          }
        }
      ],
      list: [
      ]
    }
  },
  mounted() {
    this.group_id = this.$route.params.id
    this.Request({
      url:'/qrcodes',
      method:"GET",
      params:{"group_id":this.group_id}
    }).then(response => {
      let list = response.data.data
      this.list = list
    })
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
    },
    show (index) {
      this.$Modal.info({
        title: 'User Info',
        content: `Name：${this.data6[index].name}<br>Age：${this.data6[index].age}<br>Address：${this.data6[index].address}`
      })
    },
    remove (index) {
      this.data6.splice(index, 1);
    }
  }
}
</script>
