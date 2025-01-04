import Grid from "~/models/Grid";
import type Solution from "~/models/Solution";

interface Check{
  success : boolean,
  grid : Grid,
  message: string
}

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
  
  static async getAllSolutions(grid: Grid): Promise<Solution[] | null> {
    const res = await fetch(`${useAppConfig().apiURL}/getallsolutions`, {
      method: "POST",
      body: JSON.stringify(grid),
    });

    const json = await res.json();
    if (json.error) {
      throw new Error(json.error.message);
    } else {
      
      return json ? (json as Solution[]) : null;
    }
  }

  static async checkGrid(grid:Grid): Promise<{verified:boolean,message:string}>{
    const res = await fetch(`${useAppConfig().apiURL}/check_grid`, {
      method: "POST",
      body: JSON.stringify(grid),
    });
    const json = await res.json() as Check;
    console.log(json)
    if(json){
     
      return {verified:json.success,message:json.message}
    }else{
      return {verified: false, message:"Something very bad happend :( Contact your support <3"}
    }
  }
}
