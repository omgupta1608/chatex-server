import Navbar from './Navbar.jsx';

const UserAuthFormPage = ({ children }) => {
	return (
		<div className='user-form-page'>
			<Navbar showSearchBar={false} />
			<main>{children}</main>
		</div>
	);
};

export default UserAuthFormPage;
