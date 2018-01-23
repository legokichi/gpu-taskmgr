package main

import (
	"fmt"
	"github.com/legokichi/gpu-taskmgr/io"
)

func main() {
	
}

/*
export async function main(){
	const io = new IO("./save.json", "./in", "./out");
	await io.fetch();
	const prms = [0].map((gpuId)=> (async (process)=>{
	  while(true){
		let filepath = "";
		try{
		  const opt = await io.get();
		  if(opt === null){
			console.log(gpuId, "nothing todo");
			await sleep(60*1000);
			continue;
		  }
		  console.log(gpuId, "got file", opt);
		  filepath = opt;
		  const resultDir = ""; //await process.run(gpuId, filepath);
		  await io.commit(filepath, resultDir);
		}catch(err){
		  console.log(gpuId, "ignore", filepath, err);
		  await io.ignore(filepath, err);
		}finally{
		  console.log(gpuId, "done");
		}
	  }
	})(new Prosess()).catch(console.error.bind(console, gpuId, "error")));
	const fetch = (async ()=>{ while(true){ await io.fetch(); await sleep(1000); } })();
	const save = (async ()=>{ while(true){ await io.save(); await sleep(1000); } })();
  }
*/