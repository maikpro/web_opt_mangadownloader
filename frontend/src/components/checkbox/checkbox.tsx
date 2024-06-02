import { ChangeEvent, useState } from 'react';

interface CheckboxProps {
    labelText: string;
    isChecked: boolean;
    onChange: (checked: boolean) => void;
    isDisabled?: boolean;
}

export default function Checkbox(checkboxProps: CheckboxProps) {
    const handleCheckedChange = (event: ChangeEvent<HTMLInputElement>) => {
        const { checked } = event.target;
        checkboxProps.onChange(checked);
    };

    return (
        <div>
            <input
                type="checkbox"
                name="isChecked"
                checked={checkboxProps.isChecked}
                onChange={handleCheckedChange}
                disabled={checkboxProps.isDisabled}
            />
            <label className="ml-2">{checkboxProps.labelText}</label>
        </div>
    );
}
