"127.0.0.1" => string tunnelHost;
9001 => int tunnelPort;

"127.0.0.1" => string destHost;
9002 => int destPort;

"tuntun" => string cmd;
" --listen " + tunnelHost + ":" + Std.itoa(tunnelPort) +=> cmd;
" --forward " + destHost + ":" + Std.itoa(destPort) +=> cmd;
<<< Std.system(cmd + " &") >>>;

fun void gen() {
    OscSend xmit;
    xmit.setHost(tunnelHost, tunnelPort);

    while (true) {
        xmit.startMsg("/example", "i");
        Math.random2(0,127) => xmit.addInt;
        100::ms => now;
    }
}
