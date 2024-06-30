import { toast } from 'react-toastify';

export class NotifierService {
    public static showSuccess(text: string) {
        toast.success(`${text}`);
    }

    public static showError(text: string) {
        toast.error(`${text}`);
        console.error(text);
    }
}
