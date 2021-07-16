import { Route, Redirect } from 'react-router-dom';
import useAuth from '../Hooks/useAuth';

/**
 * A wrapper for <Route> that redirects to the dashboard if user is authenticated.
 */
function AuthRoute({ children, ...rest }) {
	let { isUserAuthenticated } = useAuth();
	return (
		<Route
			{...rest}
			render={({ location }) => {
				if (!isUserAuthenticated()) return children;
				return (
					<Redirect
						to={{
							pathname: '/',
							state: { from: location },
						}}
					/>
				);
			}}
		/>
	);
}

export default AuthRoute;
