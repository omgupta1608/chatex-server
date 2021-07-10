import './App.css';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';

function App() {
	return (
		<Router>
			<Switch>
				<Route path='/login' exact>
					{/* <Login /> */}
					Login
				</Route>
				<Route path='/register' exact>
					{/* <Register /> */}
					Register
				</Route>
				<Route path='/'>Dashboard</Route>
			</Switch>
		</Router>
	);
}

export default App;
