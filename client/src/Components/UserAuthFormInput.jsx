import { useState, useCallback } from 'react';

const useInputFieldWithErr = () => {
	const [inputValue, setInputValue] = useState('');
	const [showErr, setShowErr] = useState(true);
	const [errMessage, setErrMessage] = useState('');

	const onInputChange = useCallback(e => {
		// disable error when user is typing
		setShowErr(false);

		setInputValue(e.target.value);
	}, []);

	// enable error when user moves ahead of current input
	const onInputBlur = useCallback(e => {
		setShowErr(true);

		// browser's validation message
		const validationMessage = e.target.validationMessage;
		setErrMessage(validationMessage);
	}, []);

	return { inputValue, onInputChange, onInputBlur, showErr, errMessage };
};

/**
 * @function
 * Can also pass additional input element attributes(these will override the default ones)
 * @param {Object} props
 * @param {string} props.name name of input
 * @param {string} props.label label for input
 * @param {string} props.id id of input
 */
const UserAuthFormInput = ({ id, name, label, ...additionalAttributes }) => {
	const { inputValue, onInputChange, onInputBlur, showErr, errMessage } =
		useInputFieldWithErr();

	const className = `
		input-field
		${!inputValue ? 'input-field--empty' : ''}
	`;

	return (
		<div className={className} key={id}>
			<input
				type='text'
				name={name}
				id={id}
				value={inputValue}
				onChange={onInputChange}
				onBlur={onInputBlur}
				{...additionalAttributes}
			/>
			<label htmlFor={id}>{label}</label>
			{showErr && <small>{errMessage}</small>}
		</div>
	);
};

export default UserAuthFormInput;
