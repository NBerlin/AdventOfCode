const fs = require('fs');
const file = fs.readFileSync('input', 'utf8');
//part 1 
console.log(file.split("\n").reduce((acc,change)=>acc+Number(change),0));
//part 2
const part2 = function(){
    const numbers=file.split("\n");
    const visited = {};
    let freq=0;
    while(true){
        for(const delta of numbers){
            console.log(freq)
            freq+=Number(delta);
            if(visited[freq]){
                
                return freq;
            }
            visited[freq]=true;
        }
        
    }
}
console.log(part2())
