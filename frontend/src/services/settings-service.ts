import { Settings } from '@/models/Settings';
import { toast } from 'react-toastify';
import { NotifierService } from './notifier-service';

export class SettingsService {
    public static async getSettings(): Promise<Settings | null> {
        try {
            const response = await fetch(
                `${process.env.NEXT_PUBLIC_API_URL}/settings`
            );
            if (response.ok) {
                const settings: Settings = await response.json();
                return settings;
            }
            NotifierService.showError(
                'response is not 200 by getting Settings!'
            );

            return null;
        } catch {
            NotifierService.showError('Error getting Settings!');
            return null;
        }
    }

    public static async saveSettings(settings: Settings) {
        try {
            const response = await fetch(
                `${process.env.NEXT_PUBLIC_API_URL}/settings`,
                {
                    method: 'POST',
                    body: JSON.stringify(settings),
                }
            );

            if (response.ok) {
                const savedSettings: Settings = await response.json();
                NotifierService.showSuccess('Saved Settings');
                return savedSettings;
            }
            NotifierService.showError(
                'response is not 200 by saving settings!'
            );
        } catch (error) {
            NotifierService.showError('Error saving Settings');
        }
    }

    public static async updateSettings(settings: Settings) {
        try {
            const response = await fetch(
                `${process.env.NEXT_PUBLIC_API_URL}/settings/id/${settings.id}`,
                {
                    method: 'PUT',
                    body: JSON.stringify(settings),
                }
            );

            if (response.ok) {
                const updatedSettings: Settings = await response.json();
                NotifierService.showSuccess('Updated Settings');
                return updatedSettings;
            }
            NotifierService.showError(
                'response is not 200 by updating settings!'
            );
        } catch (error) {
            NotifierService.showError('Error updating Settings');
        }
    }
}
