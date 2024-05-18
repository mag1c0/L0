export default function (t: any, e: any) {
    let n, o = null !== (n = null == t ? void 0 : t.offsetTop) && void 0 !== n ? n : 0;
    window.scrollTo({
        top: o + e,
        behavior: "smooth"
    })
}