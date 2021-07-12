import UserAuthForm from '../Components/UserAuthForm';
import UserAuthFormInput from '../Components/UserAuthFormInput';
import UserAuthFormPage from '../Components/UserAuthFormPage';

/**
 * user's email verification page
 */
const UserVerificationPage = () => {
	const onSubmit = e => {
		e.preventDefault();
		console.log(e.target.checkValidity());
	};

	return (
		<UserAuthFormPage>
			<UserAuthForm
				title='Verify your Email'
				// TODO: get email from router state
				description={`We sent a 6 digit verification code to your email - email placeholder`}
				onSubmit={onSubmit}
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
