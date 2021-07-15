import './sass/App.scss';
import { useState } from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import { RegisterPage, LoginPage, UserVerificationPage, MainPage } from './Pages';

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
		<div className={`App ${theme}`}>
			<Router>
				<Switch>
					<Route path='/login' exact>
						<LoginPage />
					</Route>
					<Route path='/register' exact>
						<RegisterPage />
					</Route>
					<Route path='/register/verify' exact>
						<UserVerificationPage />
					</Route>
					<Route path='/'><MainPage/></Route>
				</Switch>
			</Router>
		</div>
	);
}

export default App;
