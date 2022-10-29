<template>
  <div>
    <Row :gutter="20">
      <i-col :xs="12" :md="8" :lg="4" v-for="(infor, i) in inforCardData" :key="`infor-${i}`" style="height: 120px;padding-bottom: 10px;">
        <infor-card shadow :color="infor.color" :icon="infor.icon" :icon-size="36">
          <count-to :end="infor.count" count-class="count-style"/>
          <p>{{ infor.title }}</p>
        </infor-card>
      </i-col>
    </Row>
    <Row>
      <Table border :columns="report_columns" :data="report_list">
      </Table>
    </Row>
  </div>
</template>
<script>
import InforCard from '@/components/info-card'
import CountTo from '@/components/count-to'
import { ChartPie, ChartBar } from '@/components/charts'
export default {
  name: 'home',
  components: {
    InforCard,
    CountTo,
    ChartPie,
    ChartBar,
  },
  data () {
    return {
      report_list:[
        {
          "date":"2020-10-24"
        },
        {
          "date":"2020-10-23"
        },
        {
          "date":"2020-10-22"
        },
        {
          "date":"2020-10-21"
        },
        {
          "date":"2020-10-20"
        },
        {
          "date":"2020-10-19"
        },
        {
          "date":"2020-10-18"
        },
        {
          "date":"2020-10-17"
        }
      ],
      group_list:[
        {
          "title":"哈尔滨商业联盟群"
        },
        {
          "title":"哈尔滨吃喝玩乐群"
        },
      ],
      group_columns: [{
        title: '订单来源',
        align: 'center',
        children: [
          {
            title: '名称',
            key: 'title'
          },
          {
            title: '订单金额',
            key: 'date'
          }]
      }],
      report_columns: [
        {
          title: '日期',
          key: 'date'
        },
        {
          title: '当日新增用户',
          key: 'num'
        },
        {
          title: '当日新增群数',
          key: 'status'
        },
        {
          title: '当日订单金额',
          key: 'status'
        },
        {
          title: '转化率',
          key: 'status'
        },
        {
          title: '付费率',
          key: 'status'
        },
        {
          title: '累计订单金额',
          key: 'status'
        },
      ],
      columns1: [
        {
          title: '日期',
          key: 'date'
        },
        {
          title: '当日新增用户',
          key: 'num'
        },
        {
          title: '当日新增群数',
          key: 'status'
        },
        {
          title: '当日订单金额',
          key: 'status'
        },
        {
          title: '转化率',
          key: 'status'
        },
        {
          title: '付费率',
          key: 'status'
        },
        {
          title: '累计订单金额',
          key: 'status'
        },
      ],
      inforCardData: [
        { title: '当日新增用户', icon: 'md-person-add', count: 803, color: '#2d8cf0' },
        { title: '当日新增群数', icon: 'md-person-add', count: 803, color: '#2d8cf0' },
        { title: '当日订单数', icon: 'md-locate', count: 232, color: '#19be6b' },
        { title: '当日订单金额', icon: 'md-help-circle', count: 142, color: '#ff9900' },
        { title: '累计订单金额', icon: 'md-chatbubbles', count: 12, color: '#E46CBB' },
        { title: '投产比', icon: 'md-map', count: 14, color: '#9A66E4' }
      ],
      pieData: [
        { value: 335, name: '直接访问' },
        { value: 310, name: '邮件营销' },
        { value: 234, name: '联盟广告' },
        { value: 135, name: '视频广告' },
        { value: 1548, name: '搜索引擎' }
      ],
      barData: {
        Mon: 13253,
        Tue: 34235,
        Wed: 26321,
        Thu: 12340,
        Fri: 24643,
        Sat: 1322,
        Sun: 1324
      }
    }
  },
  mounted () {
  }
}
</script>
<style scoped>
.count-style{
  font-size: 50px;
}
</style>

