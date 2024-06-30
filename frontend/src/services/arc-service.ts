import { Arc } from '@/models/Arc';
import { NotifierService } from './notifier-service';

export class ArcService {
    public static async getArcs(): Promise<Arc[] | null> {
        try {
            const response = await fetch(
                `${process.env.NEXT_PUBLIC_API_URL}/arcs`
            );
            if (response.ok) {
                return await response.json();
            }
            NotifierService.showError('response is not 200 by getting Arcs!');
            return null;
        } catch {
            NotifierService.showError('Error getting Arcs');
            return null;
        }
    }
}
