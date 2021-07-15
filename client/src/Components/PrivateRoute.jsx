import { Route, Redirect } from 'react-router-dom';
import useAuth from '../Hooks/useAuth';

/**
 * A wrapper for <Route> that redirects to the login
 * screen if you're not yet authenticated.
 */
function PrivateRoute({ children, ...rest }) {
	let { isUserAuthenticated } = useAuth();
	return (
		<Route
			{...rest}
			render={({ location }) => {
				if (isUserAuthenticated()) return children;
				return (
					<Redirect
						to={{
							pathname: '/login',
							state: { from: location },
						}}
					/>
				);
			}}
		/>
	);
}

export default PrivateRoute;
