export namespace runGeneratorService {
	
	export class Location {
	    x: number;
	    y: number;
	    penalty: {[key: string]: number};
	
	    static createFrom(source: any = {}) {
	        return new Location(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.x = source["x"];
	        this.y = source["y"];
	        this.penalty = source["penalty"];
	    }
	}
	export class Activity {
	    name: string;
	    duration: number;
	    type: string;
	    fraction: string;
	    entryLocation: Location;
	    exitLocation: Location;
	    fastTravel: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Activity(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.duration = source["duration"];
	        this.type = source["type"];
	        this.fraction = source["fraction"];
	        this.entryLocation = this.convertValues(source["entryLocation"], Location);
	        this.exitLocation = this.convertValues(source["exitLocation"], Location);
	        this.fastTravel = source["fastTravel"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	

}

