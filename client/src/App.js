import './App.css';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import LoginPage from './Pages/LoginPage/LoginPage';
import RegisterPage from './Pages/RegisterPage/RegisterPage';

function App() {
	return (
		<Router>
			<Switch>
				<Route path='/login' exact>
					<LoginPage />
				</Route>
				<Route path='/register' exact>
					<RegisterPage />
				</Route>
				<Route path='/'>Dashboard</Route>
			</Switch>
		</Router>
	);
}

export default App;
