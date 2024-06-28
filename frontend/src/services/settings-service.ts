import { Settings } from '@/models/Settings';

export class SettingsService {
    public static async getSettings(): Promise<Settings> {
        try {
            const response = await fetch(
                `${process.env.NEXT_PUBLIC_API_URL}/settings`
            );
            if (response.ok) {
                const settings: Settings = await response.json();
                return settings;
            }
            throw Error('response is not 200 by getting Settings!');
        } catch {
            throw Error('Error getting Settings');
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
                return savedSettings;
            }
            throw Error('response is not 200 by saving settings!');
        } catch (error) {
            throw Error('Error saving Settings');
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
                return updatedSettings;
            }
            throw Error('response is not 200 by updating settings!');
        } catch (error) {
            throw Error('Error updating Settings');
        }
    }
}
