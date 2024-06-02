import { Arc } from '@/models/Arc';

export class ArcService {
    public static async getArcs(): Promise<Arc[]> {
        const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/arcs`);
        return await response.json();
    }
}
