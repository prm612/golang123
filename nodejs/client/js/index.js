import Vue 		 from 'vue';
import VueRouter from 'vue-router';
import iView     from 'iView';
import axios     from 'axios';
import 'iview/dist/styles/iview.css';

import App		 	  from './Index/App';
import ArticleEdit    from './Index/ArticleEdit';
import ArticleList    from './Index/ArticleList';
import CategoryManage from './Index/CategoryManage';

const routes = [
	{
		path: '/category/manage',
		component: CategoryManage
	},
	{
		path: '/article',
		component: ArticleList
	},
	{
		path: '/article/add',
		component: ArticleEdit
	},
	{
		path: '/article/edit/:id',
		component: ArticleEdit
	}
];

Vue.use(VueRouter);
Vue.use(iView);
Vue.prototype.$http = axios;

const router = new VueRouter({
	routes
})

new Vue({
	el: '#app',
	router,
	render: h => h(App)
})