import Grid from "~/models/Grid";
import type Solution from "~/models/Solution";

export default abstract class {
  static async solve(grid: Grid): Promise<Grid | null> {
    const res = await fetch(`${useAppConfig().apiURL}/solve`, {
      method: "POST",
      body: JSON.stringify(grid),
    });

    const json = await res.json();
    if (json.error) {
      throw new Error(json.error.message);
    } else {
      return json ? new Grid(json.grid as number[][]) : null;
    }
  }

  static async getOneSolution(grid: Grid): Promise<Solution | null> {
    const res = await fetch(`${useAppConfig().apiURL}/getsolution`, {
      method: "POST",
      body: JSON.stringify(grid),
    });

    const json = await res.json();
    if (json.error) {
      throw new Error(json.error.message);
    } else {
      
      return json ? (json as Solution) : null;
    }
  }
}
