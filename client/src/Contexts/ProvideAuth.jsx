import useProvideAuth from '../Hooks/useProvideAuth';
import { createContext } from 'react';

export const authContext = createContext();

const ProvideAuth = ({ children }) => {
	return (
		<authContext.Provider value={useProvideAuth()}>
			{children}
		</authContext.Provider>
	);
};

export default ProvideAuth;
