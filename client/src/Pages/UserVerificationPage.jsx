import UserAuthForm from '../Components/UserAuthForm';
import UserAuthFormInput from '../Components/UserAuthFormInput';
import UserAuthFormPage from '../Components/UserAuthFormPage';
import { Redirect } from 'react-router-dom';
import useAuth from '../Hooks/useAuth';

/**
 * user's email verification page
 */
const UserVerificationPage = () => {
	const {
		state: { errorMsg, isLoading, user },
		verifyUserAndLogin,
		isUserRegistered,
	} = useAuth();

	const onSubmit = async e => {
		e.preventDefault();
		const formElement = e.target;
		if (!formElement.checkValidity()) return;

		const formData = Object.fromEntries(new FormData(formElement));
		const reqData = {
			verification_code: formData.verification_code,
			uid: user.uid,
		};

		verifyUserAndLogin(reqData);
	};

	// redirect to register page if no user
	if (!isUserRegistered()) return <Redirect to='/register' />;
	return (
		<UserAuthFormPage>
			<UserAuthForm
				title='Verify your Email'
				// TODO: get email from router state
				description={`We sent a 6 digit verification code to your email${
					user?.email ? ` - ${user.email}` : ''
				}`}
				onSubmit={onSubmit}
				isLoading={isLoading}
				submitErrMsg={errorMsg}
			>
				<UserAuthFormInput
					name='verification_code'
					id='user-verification-code'
					label='Verification Code'
					type='text'
					minLength={6}
					maxLength={6}
					required={true}
					autoFocus={true}
				/>
			</UserAuthForm>
		</UserAuthFormPage>
	);
};

export default UserVerificationPage;
