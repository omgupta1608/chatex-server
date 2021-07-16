import './sass/App.scss';
import { useState } from 'react';
import { BrowserRouter as Router, Switch } from 'react-router-dom';
import { RegisterPage, LoginPage, UserVerificationPage, MainPage } from './Pages';
import ProvideAuth from './Contexts/ProvideAuth';
import AuthRoute from './Components/AuthRoute';
import PrivateRoute from './Components/PrivateRoute';

// ! temperory solution. change this to a contexts
const themes = {
	LIGHT: 'theme-light',
	DARK: 'theme-dark',
};

const useTheme = initialTheme => {
	const [theme, setTheme] = useState(initialTheme);
	const toggleTheme = () =>
		setTheme(theme => (theme === themes.LIGHT ? themes.DARK : themes.LIGHT));

	return { theme, setTheme, toggleTheme };
};

function App() {
	const { theme } = useTheme(themes.DARK);

	return (
		<ProvideAuth>
			<div className={`App ${theme}`}>
				<Router>
					<Switch>
						<AuthRoute path='/login' exact>
							<LoginPage />
						</AuthRoute>
						<AuthRoute path='/register' exact>
							<RegisterPage />
						</AuthRoute>
						<AuthRoute path='/register/verify' exact>
							<UserVerificationPage />
						</AuthRoute>
						<PrivateRoute path='/'><MainPage /></PrivateRoute>
					</Switch>
				</Router>
			</div>
		</ProvideAuth>
	);
}

export default App;
