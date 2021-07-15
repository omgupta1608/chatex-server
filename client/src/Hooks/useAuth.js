import { useContext } from 'react';
import { authContext } from '../Contexts/ProvideAuth';

const useAuth = () => {
	return useContext(authContext);
};

export default useAuth;
