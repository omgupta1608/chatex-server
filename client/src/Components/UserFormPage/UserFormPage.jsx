import Navbar from '../Navbar/Navbar.jsx';
import './UserFormPage.scss';

const UserFormPage = ({ children }) => {
	return (
		<div className='user-form-page'>
			<Navbar showSearchBar={false} />
			<main>{children}</main>
		</div>
	);
};

export default UserFormPage;
