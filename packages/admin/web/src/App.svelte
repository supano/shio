<script>
import createRouter from '@spaceavocado/svelte-router';
import RouterView from '@spaceavocado/svelte-router/component/view';

// View components
import HomePage from './pages/Home.svelte';
import LoginPage from './pages/Login.svelte';
import Page404 from './pages/404.svelte';  

createRouter({
  routes: [
    { 
      path: '/',
      name: 'HOME',
      component: HomePage,
	},
	    { 
      path: '/login',
      name: 'Login',
      component: LoginPage,
    },
    {
      path: '*',
      component: Page404,
    },
  ],
    mode: 'HISTORY',
});
</script>

<RouterView />