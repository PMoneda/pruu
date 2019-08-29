import TabItem from "./tabItem";

export default class TabList {
    private tabs: TabItem[]=[]
    private currentSelectedTab:number=0;


    addTabItem(tab:TabItem): void {
        this.tabs.push(tab);
    }

    removeTabItem(index:number):void{
        this.tabs.splice(index,1);
        if(this.currentSelectedTab >= this.tabs.length){
            this.setSelectedTab(this.tabs.length - (this.currentSelectedTab-this.tabs.length));
        }
    }

    getHeaders():React.ReactNodeArray{
        return this.tabs;
    }

    getCurrentContent():React.ReactNode{
        return this.tabs[this.currentSelectedTab].body();
    }

    setSelectedTab(index:number):void {
        if(index < 0 ){
            this.currentSelectedTab = this.tabs.length - 1;
        }else if(index >= this.tabs.length){
            this.currentSelectedTab = 0;
        }else{
            this.currentSelectedTab = index;
        }        
    }

    getLastTabIndex():number {
        if(this.tabs.length === 0){
            return 0;
        }
        return this.tabs.length-1;
    }
}